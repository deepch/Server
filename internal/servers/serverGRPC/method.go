/*
	Autor Andrey Semochkin
*/

package serverGRPC

import (
	"context"
	"encoding/json"

	"github.com/deepch/Server/internal/device"

	"github.com/deepch/Server/internal/memory"

	"github.com/deepch/Server/internal/constant"

	pb "github.com/deepch/Server/api/proto/server"
)

/*

	GRPC server method set

*/

//serverGRPC basic service structure
type serverGRPC struct {
	pb.UnimplementedGreeterServer
	service *ServiceGRPC
}

//InformationPlatformSummary - method gives information about the system
func (s *serverGRPC) InformationPlatformSummary(ctx context.Context, _ *pb.Empty) (*pb.CodedReply, error) {

	//Authorization token verification
	err := s.service.authentication(ctx)

	if err != nil {

		return nil, err

	}

	return &pb.CodedReply{Message: constant.ConstNone}, nil

}

//InformationDeviceStatistic - method returns device statistics
func (s *serverGRPC) InformationDeviceStatistic(ctx context.Context, _ *pb.DeviceUniqID) (*pb.CodedReply, error) {

	//Authorization token verification
	err := s.service.authentication(ctx)

	if err != nil {

		return nil, err

	}

	return &pb.CodedReply{Message: constant.ConstNone}, nil
}

//InformationDeviceStatus - method returns device operation status
func (s *serverGRPC) InformationDeviceStatus(ctx context.Context, in *pb.DeviceUniqID) (*pb.MapBoolReply, error) {

	//Authorization token verification
	err := s.service.authentication(ctx)

	if err != nil {
		return nil, err
	}

	tmp, err := memory.TnxMemory.SelectDevice(in.DeviceUniqID).SelectStatus()

	if err != nil {
		return nil, err
	}

	return &pb.MapBoolReply{Channels: tmp}, nil
}

//InformationDeviceSummary - the function gets all device parameters
func (s *serverGRPC) InformationDeviceSummary(ctx context.Context, in *pb.DeviceUniqID) (*pb.ConfigurationsDeviceRequest, error) {

	//Authorization token verification
	err := s.service.authentication(ctx)

	if err != nil {
		return nil, err
	}

	tmp, err := memory.TnxMemory.SelectDevice(in.DeviceUniqID).SelectSummary()

	if err != nil {
		return nil, err
	}

	if tmp != nil {
		res := &pb.ConfigurationsDeviceRequest{
			Name:     tmp.Name,
			OnvifURL: tmp.OnvifURL,
			Channels: make(map[string]*pb.DeviceChannel),
		}

		if tmp.Channels != nil {

			for s2, c := range tmp.Channels {

				res.Channels[s2] = &pb.DeviceChannel{
					Name:        c.Name,
					DataURL:     c.DataURL,
					SnapshotURL: c.SnapshotURL,
					Audio:       c.Audio,
					Record:      c.Record,
					VOD:         c.VOD,
					Debug:       c.Debug,
					DebugRaw:    c.DebugRaw,
					RTPProxy:    c.RTPProxy,
					Status:      c.Status,
					Bitrate:     c.Bitrate,
					RecordTime:  c.RecordTime,
				}

			}

		}

		return res, nil

	}

	return nil, constant.ErrDeviceNotFound

}

//InformationDevicesSummary - all parameters of all devices
func (s *serverGRPC) InformationDevicesSummary(ctx context.Context, _ *pb.Empty) (*pb.DevicesSummaryReply, error) {

	//Authorization token verification
	err := s.service.authentication(ctx)

	if err != nil {
		return nil, err
	}

	res := pb.DevicesSummaryReply{Encoders: make(map[string]*pb.ConfigurationsDeviceRequest)}

	tmp := memory.TnxMemory.SelectSummary()

	for i, i2 := range tmp {

		if i2 != nil {

			res.Encoders[i] = &pb.ConfigurationsDeviceRequest{
				Name:     i2.Name,
				OnvifURL: i2.OnvifURL,
				Channels: make(map[string]*pb.DeviceChannel),
			}

			if i2.Channels != nil {

				for s2, c := range i2.Channels {

					res.Encoders[i].Channels[s2] = &pb.DeviceChannel{
						Name:        c.Name,
						DataURL:     c.DataURL,
						SnapshotURL: c.SnapshotURL,
						Audio:       c.Audio,
						Record:      c.Record,
						VOD:         c.VOD,
						Debug:       c.Debug,
						DebugRaw:    c.DebugRaw,
						RTPProxy:    c.RTPProxy,
						Status:      c.Status,
						Bitrate:     c.Bitrate,
						RecordTime:  c.RecordTime,
					}

				}

			}

		}

	}

	return &res, nil

}

//InformationPlatformVersion - method display system version
func (s *serverGRPC) InformationPlatformVersion(_ context.Context, _ *pb.Empty) (*pb.CodedReply, error) {

	return &pb.CodedReply{Message: constant.ConstVersion}, nil
}

//InformationPlatformPing - response to a ping request (pong)
func (s *serverGRPC) InformationPlatformPing(_ context.Context, _ *pb.Empty) (*pb.CodedReply, error) {

	return &pb.CodedReply{Message: constant.ConstPong}, nil
}

//ConfigurationsDeviceAddition - adding a new device
func (s *serverGRPC) ConfigurationsDeviceAddition(ctx context.Context, in *pb.ConfigurationsDeviceRequest) (*pb.CodedReply, error) {

	//Authorization token verification
	err := s.service.authentication(ctx)

	if err != nil {
		return nil, err
	}

	res, err := json.Marshal(in)

	if err != nil {
		return nil, err
	}

	var tmp device.Device

	err = json.Unmarshal(res, &tmp)

	if err != nil {
		return nil, err
	}

	err = memory.TnxMemory.Put(in.DeviceUniqID, &tmp)

	if err != nil {

		return nil, err

	}

	return &pb.CodedReply{Message: constant.ConstSuccess}, nil
}

//ConfigurationsDeviceEditing - device editing
func (s *serverGRPC) ConfigurationsDeviceEditing(ctx context.Context, in *pb.ConfigurationsDeviceRequest) (*pb.CodedReply, error) {

	//Authorization token verification
	err := s.service.authentication(ctx)

	if err != nil {
		return nil, err
	}

	res, err := json.Marshal(in)

	if err != nil {
		return nil, err
	}

	var tmp device.Device

	err = json.Unmarshal(res, &tmp)

	if err != nil {
		return nil, err
	}

	err = memory.TnxMemory.Set(in.DeviceUniqID, &tmp)

	if err != nil {
		return nil, err
	}

	return &pb.CodedReply{Message: constant.ConstSuccess}, nil

}

//ConfigurationsDeviceDeletion - device removal
func (s *serverGRPC) ConfigurationsDeviceDeletion(ctx context.Context, in *pb.DeviceUniqID) (*pb.CodedReply, error) {

	//Authorization token verification
	err := s.service.authentication(ctx)

	if err != nil {
		return nil, err
	}

	err = memory.TnxMemory.Delete(in.DeviceUniqID)

	if err != nil {
		return nil, err
	}

	return &pb.CodedReply{Message: constant.ConstSuccess}, nil

}

//ConfigurationsDeviceReconnection - forced reconnection to the device
func (s *serverGRPC) ConfigurationsDeviceReconnection(ctx context.Context, _ *pb.DeviceUniqID) (*pb.CodedReply, error) {

	//Authorization token verification
	err := s.service.authentication(ctx)

	if err != nil {
		return nil, err
	}

	return &pb.CodedReply{Message: constant.ConstSuccess}, nil

}
