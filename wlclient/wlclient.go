// Package wlclient implements a wayland-client like api
package wlclient

import "github.com/neurlang/wayland/wl"
import "github.com/neurlang/wayland/xdg"
import "github.com/neurlang/wayland/unstable"

func DisplayDispatch(d *wl.Display) error {
	return d.Context().Run()
}
func PointerSetUserData(p *wl.Pointer, data interface{}) {
}

func SurfaceSetUserData(p *wl.Surface, data interface{}) {
	p.UserData = data
}

type PointerListener interface {
	wl.PointerEnterHandler
	wl.PointerLeaveHandler
	wl.PointerMotionHandler
	wl.PointerButtonHandler
	wl.PointerAxisHandler
	wl.PointerFrameHandler
	wl.PointerAxisSourceHandler
	wl.PointerAxisStopHandler
	wl.PointerAxisDiscreteHandler
}

func PointerAddListener(p *wl.Pointer, h PointerListener) {
	p.AddEnterHandler(h)
	p.AddLeaveHandler(h)
	p.AddMotionHandler(h)
	p.AddButtonHandler(h)
	p.AddAxisHandler(h)
	p.AddFrameHandler(h)
	p.AddAxisSourceHandler(h)
	p.AddAxisStopHandler(h)
	p.AddAxisDiscreteHandler(h)

}
func PointerDestroy(p *wl.Pointer) {
	//p.Destroy()
	p.Unregister()
}
func ShmDestroy(p *wl.Shm) {
	//p.Destroy()
	p.Unregister()
}
func RegistryDestroy(p *wl.Registry) {
	//p.Destroy()
	p.Unregister()
}
func BufferAddListener(b *wl.Buffer, data wl.BufferReleaseHandler) {
	b.AddReleaseHandler(data)
}
func CallbackAddListener(c *wl.Callback, data wl.CallbackDoneHandler) {
	c.AddDoneHandler(data)
}

type OutputListener interface {
	wl.OutputGeometryHandler
	wl.OutputModeHandler
	wl.OutputDoneHandler
	wl.OutputScaleHandler
}

func OutputAddListener(o *wl.Output, h OutputListener) {
	o.AddGeometryHandler(h)
	o.AddModeHandler(h)
	o.AddDoneHandler(h)
	o.AddScaleHandler(h)
}

type SeatListener interface {
	wl.SeatCapabilitiesHandler
	wl.SeatNameHandler
}

func SeatDestroy(p *wl.Seat) {
	//p.Destroy()
	p.Unregister()
}
func SeatAddListener(s *wl.Seat, data SeatListener) {
	s.AddCapabilitiesHandler(data)
	s.AddNameHandler(data)
}

type RegistryListener interface {
	wl.RegistryGlobalHandler
	wl.RegistryGlobalRemoveHandler
}

func RegistryAddListener(r *wl.Registry, data RegistryListener) {
	r.AddGlobalHandler(data)
	r.AddGlobalRemoveHandler(data)
}
func KeyboardSetUserData(*wl.Keyboard, interface{}) {
}

type KeyboardListener interface {
	wl.KeyboardKeymapHandler
	wl.KeyboardEnterHandler
	wl.KeyboardLeaveHandler
	wl.KeyboardKeyHandler
	wl.KeyboardModifiersHandler
	wl.KeyboardRepeatInfoHandler
}

func KeyboardAddListener(kb *wl.Keyboard, l KeyboardListener) {
	kb.AddKeymapHandler(l)
	kb.AddEnterHandler(l)
	kb.AddLeaveHandler(l)
	kb.AddKeyHandler(l)
	kb.AddModifiersHandler(l)
	kb.AddRepeatInfoHandler(l)
}
func KeyboardDestroy(p *wl.Keyboard) {
	//p.Destroy()
	p.Unregister()
}
func TouchSetUserData(*wl.Touch, interface{}) {
}

type TouchListener interface {
	wl.TouchDownHandler
	wl.TouchUpHandler
	wl.TouchMotionHandler
	wl.TouchFrameHandler
	wl.TouchCancelHandler
	wl.TouchShapeHandler
	wl.TouchOrientationHandler
}

func TouchAddListener(to *wl.Touch, tl TouchListener) {
	to.AddDownHandler(tl)
	to.AddUpHandler(tl)
	to.AddMotionHandler(tl)
	to.AddFrameHandler(tl)
	to.AddCancelHandler(tl)
	to.AddShapeHandler(tl)
	to.AddOrientationHandler(tl)
}
func TouchDestroy(p *wl.Touch) {
	//p.Destroy()
	p.Unregister()
}

type SurfaceEnterLeave struct {
	surface   *wl.Surface
	callbacks [2]func(*wl.Surface, *wl.Output)
}

func (el *SurfaceEnterLeave) HandleSurfaceEnter(en wl.SurfaceEnterEvent) {
	el.callbacks[0](el.surface, en.Output)
}

func (el *SurfaceEnterLeave) HandleSurfaceLeave(le wl.SurfaceLeaveEvent) {
	el.callbacks[1](el.surface, le.Output)
}

func SurfaceAddListener(
	s *wl.Surface,
	enter func(*wl.Surface, *wl.Output),
	leave func(*wl.Surface, *wl.Output),
) {
	el := &SurfaceEnterLeave{surface: s, callbacks: [2]func(*wl.Surface, *wl.Output){enter, leave}}
	s.AddEnterHandler(el)
	s.AddLeaveHandler(el)
}

func ShmAddListener(p *wl.Shm, data wl.ShmFormatHandler) {
	p.AddFormatHandler(data)
}
func RegionDestroy(p *wl.Region) {
	p.Destroy()
	p.Unregister()
}
func CallbackDestroy(p *wl.Callback) {
	//p.Destroy()
	p.Unregister()
}
func SubsurfaceDestroy(p *wl.Subsurface) {
	p.Destroy()
	p.Unregister()
}
func DataDeviceDestroy(p *wl.DataDevice) {
	//p.Destroy()
	p.Unregister()
}
func DataDeviceManagerDestroy(d *wl.DataDeviceManager) {
	//d.Destroy()
	d.Unregister()
}

type DataDeviceListener interface {
	wl.DataDeviceDataOfferHandler
	wl.DataDeviceEnterHandler
	wl.DataDeviceLeaveHandler
	wl.DataDeviceMotionHandler
	wl.DataDeviceDropHandler
	wl.DataDeviceSelectionHandler
}

func DataDeviceAddListener(p *wl.DataDevice, h DataDeviceListener) {
	p.AddDataOfferHandler(h)
	p.AddEnterHandler(h)
	p.AddLeaveHandler(h)
	p.AddMotionHandler(h)
	p.AddDropHandler(h)
	p.AddSelectionHandler(h)
}

type DataOfferListener interface {
	wl.DataOfferOfferHandler
	wl.DataOfferSourceActionsHandler
	wl.DataOfferActionHandler
}

func DataOfferDestroy(p *wl.DataOffer) {
	//d.Destroy()
	p.Unregister()
}
func DataOfferAddListener(p *wl.DataOffer, h DataOfferListener) {
	p.AddOfferHandler(h)
	p.AddSourceActionsHandler(h)
	p.AddActionHandler(h)
}

type DataSourceListener interface {
	wl.DataSourceTargetHandler
	wl.DataSourceSendHandler
	wl.DataSourceCancelledHandler
	wl.DataSourceDndDropPerformedHandler
	wl.DataSourceDndFinishedHandler
	wl.DataSourceActionHandler
}

func DataSourceAddListener(p *wl.DataSource, h DataSourceListener) {
	p.AddTargetHandler(h)
	p.AddSendHandler(h)
	p.AddCancelledHandler(h)
	p.AddDndDropPerformedHandler(h)
	p.AddDndFinishedHandler(h)
	p.AddActionHandler(h)
}

func DataSourceRemoveListener(p *wl.DataSource, h DataSourceListener) {
	p.RemoveTargetHandler(h)
	p.RemoveSendHandler(h)
	p.RemoveCancelledHandler(h)
	p.RemoveDndDropPerformedHandler(h)
	p.RemoveDndFinishedHandler(h)
	p.RemoveActionHandler(h)
}

func RegistryBindCompositorInterface(r *wl.Registry, name uint32, version uint32) *wl.Compositor {
	c := wl.NewCompositor(r.Ctx)
	_ = r.Bind(name, "wl_compositor", version, c)
	return c
}

func RegistryBindShmInterface(r *wl.Registry, name uint32, version uint32) *wl.Shm {
	s := wl.NewShm(r.Ctx)
	_ = r.Bind(name, "wl_shm", version, s)
	return s
}

func RegistryBindDataDeviceManagerInterface(
	r *wl.Registry,
	name uint32,
	version uint32,
) *wl.DataDeviceManager {
	d := wl.NewDataDeviceManager(r.Ctx)
	_ = r.Bind(name, "wl_data_device_manager", version, d)
	return d
}

func RegistryBindOutputInterface(r *wl.Registry, name uint32, version uint32) *wl.Output {
	d := wl.NewOutput(r.Ctx)
	_ = r.Bind(name, "wl_output", version, d)
	return d
}

func RegistryBindSeatInterface(r *wl.Registry, name uint32, version uint32) *wl.Seat {
	d := wl.NewSeat(r.Ctx)
	_ = r.Bind(name, "wl_seat", version, d)
	return d
}

func RegistryBindWmBaseInterface(r *wl.Registry, name uint32, version uint32) *xdg.WmBase {
	d := xdg.NewShell(r.Ctx)
	_ = r.Bind(name, "xdg_wm_base", version, d)
	return d
}

func RegistryBindUnstableInterface(
	r *wl.Registry,
	name uint32,
	iface string,
	version uint32,
) wl.Proxy {
	function := unstable.GetNewFunc(iface)
	if function == nil {
		return nil
	}
	d := function(r.Ctx)
	_ = r.Bind(name, iface, version, d)
	return d
}

func DisplayConnect(name []byte) (*wl.Display, error) {
	return wl.Connect(string(name))
}
func DisplayGetRegistry(d *wl.Display) (*wl.Registry, error) {
	return d.GetRegistry()
}

func DisplayRun(d *wl.Display) error {
	return d.Context().Run()
}
func DisplayRoundtrip(d *wl.Display) error {
	cb, err := d.Sync()
	if err != nil {
		return err
	}
	err = d.Context().RunTill(cb)
	return err
}
func DisplayDisconnect(display *wl.Display) {
	display.Context().Close()
}
