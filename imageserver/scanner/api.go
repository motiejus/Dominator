package scanner

import (
	"github.com/Symantec/Dominator/lib/image"
	"github.com/Symantec/Dominator/lib/objectserver"
	"io"
	"log"
	"sync"
)

// TODO: the types should probably be moved into a separate package, leaving
//       behind the scanner code.

type Object struct {
	length uint64
}

type notifiers map[<-chan string]chan<- string

type ImageDataBase struct {
	sync.RWMutex
	// Protected by lock.
	baseDir         string
	imageMap        map[string]*image.Image
	addNotifiers    notifiers
	deleteNotifiers notifiers
	// Unprotected by lock.
	objectServer objectserver.ObjectServer
	logger       *log.Logger
}

func LoadImageDataBase(baseDir string, objSrv objectserver.ObjectServer,
	logger *log.Logger) (*ImageDataBase, error) {
	return loadImageDataBase(baseDir, objSrv, logger)
}

func (imdb *ImageDataBase) WriteHtml(writer io.Writer) {
	imdb.writeHtml(writer)
}

func (imdb *ImageDataBase) AddImage(image *image.Image, name string) error {
	return imdb.addImage(image, name)
}

func (imdb *ImageDataBase) CheckImage(name string) bool {
	return imdb.checkImage(name)
}

func (imdb *ImageDataBase) DeleteImage(name string) error {
	return imdb.deleteImage(name)
}

func (imdb *ImageDataBase) GetImage(name string) *image.Image {
	return imdb.getImage(name)
}

func (imdb *ImageDataBase) ListImages() []string {
	return imdb.listImages()
}

func (imdb *ImageDataBase) CountImages() uint {
	return imdb.countImages()
}

func (imdb *ImageDataBase) ObjectServer() objectserver.ObjectServer {
	return imdb.objectServer
}

func (imdb *ImageDataBase) RegisterAddNotifier() <-chan string {
	return imdb.registerAddNotifier()
}

func (imdb *ImageDataBase) UnregisterAddNotifier(channel <-chan string) {
	imdb.unregisterAddNotifier(channel)
}

func (imdb *ImageDataBase) RegisterDeleteNotifier() <-chan string {
	return imdb.registerDeleteNotifier()
}

func (imdb *ImageDataBase) UnregisterDeleteNotifier(channel <-chan string) {
	imdb.unregisterDeleteNotifier(channel)
}
