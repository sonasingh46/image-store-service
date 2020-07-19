package miniostore

import (
	"github.com/minio/minio-go/v6"
	"github.com/pkg/errors"
)

// MinioStore is a minio object storage
// Please refer https://min.io/ to learn more about minio
type MinioStore struct {
	// Client minio client that can help communicate to minio APIs.
	Client *minio.Client
}

// NewMinioStore returns a new instance of minio store
func NewMinioStore(msc *MinioStoreConfig)(*MinioStore,error)  {
	endpoint := msc.HostIP+":"+msc.Port
	accessKeyID := msc.AccessKeyID
	secretAccessKey := msc.Secret
	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, msc.UseSSL)
	if err != nil {
		return nil,errors.Errorf("failed to create minio client:{%s}",err.Error())
	}
	return &MinioStore{
		Client:minioClient,
	},nil
}

// MinioStoreConfig is the config to bring a new minio store instance
type MinioStoreConfig struct {
	// HostIP is the minio object store host ip.
	HostIP string
	// Port is the port on which minio object store listens.
	Port string
	// AccessKeyID for minio object store.
	AccessKeyID string
	// Secret for minio object store.
	Secret string
	// UseSSL -- set this to true to use SSL in the transport layer.
	UseSSL bool
}

// NewMinioStoreConfig returns an empty MinioStoreConfig isntance.
func NewMinioStoreConfig()*MinioStoreConfig  {
	return &MinioStoreConfig{}
}

// WithHostIP sets the host ip in the MinioStoreConfig.
func (msc *MinioStoreConfig)WithHostIP(hostIP string)*MinioStoreConfig  {
	msc.HostIP=hostIP
	return msc
}

// WithHostPort sets the host port in the MinioStoreConfig.
func (msc *MinioStoreConfig)WithHostPort(port string)*MinioStoreConfig  {
	msc.Port=port
	return msc
}

// WithAccessKeyID sets the access id in the MinioStoreConfig.
func (msc *MinioStoreConfig)WithAccessKeyID(accessKeyID string)*MinioStoreConfig  {
	msc.AccessKeyID=accessKeyID
	return msc
}

// WithSecret sets the secret in the MinioStoreConfig.
func (msc *MinioStoreConfig)WithSecret(secret string)*MinioStoreConfig  {
	msc.Secret=secret
	return msc
}

// WithSSL sets ssl flag to true.
func (msc *MinioStoreConfig)WithSSL()*MinioStoreConfig  {
	msc.UseSSL=true
	return msc
}