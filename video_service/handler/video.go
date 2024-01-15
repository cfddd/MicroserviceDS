package handler

import video_pb "video_service/server"

type VideoService struct {
	video_pb.UnimplementedVideoServiceServer
}
