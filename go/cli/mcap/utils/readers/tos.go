package readers

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/volcengine/ve-tos-golang-sdk/v2/tos"
)

// Automatically register TOS reader when imported.
func init() {
	RegisterReader("tos", newTOSReader)
}

// Factory for TOS readers (called by registry).
func newTOSReader(ctx context.Context, bucket, path string) (func() error, io.ReadSeekCloser, error) {
	// Get credentials from environment variables
	accessKey := os.Getenv("TOS_ACCESS_KEY")
	secretKey := os.Getenv("TOS_SECRET_KEY")
	endpoint := os.Getenv("TOS_ENDPOINT")
	region := os.Getenv("TOS_REGION")

	// Validate required configuration
	if accessKey == "" || secretKey == "" {
		return func() error { return nil }, nil, fmt.Errorf("TOS_ACCESS_KEY and TOS_SECRET_KEY environment variables must be set")
	}
	if endpoint == "" {
		return func() error { return nil }, nil, fmt.Errorf("TOS_ENDPOINT environment variable must be set (e.g., tos-cn-beijing.volces.com)")
	}
	if region == "" {
		return func() error { return nil }, nil, fmt.Errorf("TOS_REGION environment variable must be set (e.g., cn-beijing)")
	}

	// Create TOS client
	credential := tos.NewStaticCredentials(accessKey, secretKey)
	client, err := tos.NewClientV2(endpoint, tos.WithCredentials(credential), tos.WithRegion(region))
	if err != nil {
		return func() error { return nil }, nil, fmt.Errorf("failed to create TOS client: %w", err)
	}

	rs, err := NewTOSReadSeekCloser(ctx, client, bucket, path)
	if err != nil {
		client.Close()
		return func() error { return nil }, nil, fmt.Errorf("failed to create TOS reader: %w", err)
	}

	// Wrap client.Close() to match func() error signature
	closeFunc := func() error {
		client.Close()
		return nil
	}

	return closeFunc, rs, nil
}

// TOSReadSeekCloser implements io.ReadSeekCloser for TOS objects.
type TOSReadSeekCloser struct {
	ctx    context.Context
	client *tos.ClientV2
	bucket string
	key    string
	reader io.ReadCloser
	size   int64
	offset int64
}

// NewTOSReadSeekCloser creates a seekable reader for a TOS object.
func NewTOSReadSeekCloser(ctx context.Context, client *tos.ClientV2, bucket, key string) (*TOSReadSeekCloser, error) {
	// Get object metadata to retrieve size
	head, err := client.HeadObjectV2(ctx, &tos.HeadObjectV2Input{
		Bucket: bucket,
		Key:    key,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to head TOS object: %w", err)
	}

	// Get the object
	resp, err := client.GetObjectV2(ctx, &tos.GetObjectV2Input{
		Bucket: bucket,
		Key:    key,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to open TOS object: %w", err)
	}

	return &TOSReadSeekCloser{
		ctx:    ctx,
		client: client,
		bucket: bucket,
		key:    key,
		reader: resp.Content,
		size:   head.ContentLength,
		offset: 0,
	}, nil
}

func (r *TOSReadSeekCloser) Read(p []byte) (int, error) {
	n, err := r.reader.Read(p)
	r.offset += int64(n)
	return n, err
}

func (r *TOSReadSeekCloser) Seek(offset int64, whence int) (int64, error) {
	var target int64
	switch whence {
	case io.SeekStart:
		target = offset
	case io.SeekCurrent:
		target = r.offset + offset
	case io.SeekEnd:
		target = r.size + offset
	default:
		return 0, fmt.Errorf("invalid whence: %d", whence)
	}

	if target == r.offset {
		return target, nil
	}
	if target < 0 || target > r.size {
		return 0, fmt.Errorf("seek out of bounds: %d", target)
	}

	_ = r.reader.Close()

	// Use Range header to read from target offset to end of file
	// Format: "bytes=start-" means read from start to end
	rangeHeader := fmt.Sprintf("bytes=%d-", target)
	resp, err := r.client.GetObjectV2(r.ctx, &tos.GetObjectV2Input{
		Bucket: r.bucket,
		Key:    r.key,
		Range:  rangeHeader,
	})
	if err != nil {
		return 0, fmt.Errorf("failed to reopen TOS object: %w", err)
	}

	r.reader = resp.Content
	r.offset = target
	return target, nil
}

func (r *TOSReadSeekCloser) Close() error {
	return r.reader.Close()
}
