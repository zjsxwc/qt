package nfc

//#include "nfc.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QNearFieldShareTarget struct {
	core.QObject
}

type QNearFieldShareTarget_ITF interface {
	core.QObject_ITF
	QNearFieldShareTarget_PTR() *QNearFieldShareTarget
}

func PointerFromQNearFieldShareTarget(ptr QNearFieldShareTarget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNearFieldShareTarget_PTR().Pointer()
	}
	return nil
}

func NewQNearFieldShareTargetFromPointer(ptr unsafe.Pointer) *QNearFieldShareTarget {
	var n = new(QNearFieldShareTarget)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QNearFieldShareTarget_") {
		n.SetObjectName("QNearFieldShareTarget_" + qt.Identifier())
	}
	return n
}

func (ptr *QNearFieldShareTarget) QNearFieldShareTarget_PTR() *QNearFieldShareTarget {
	return ptr
}

func (ptr *QNearFieldShareTarget) Cancel() {
	defer qt.Recovering("QNearFieldShareTarget::cancel")

	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_Cancel(ptr.Pointer())
	}
}

func (ptr *QNearFieldShareTarget) ConnectError(f func(error QNearFieldShareManager__ShareError)) {
	defer qt.Recovering("connect QNearFieldShareTarget::error")

	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_ConnectError(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error", f)
	}
}

func (ptr *QNearFieldShareTarget) DisconnectError() {
	defer qt.Recovering("disconnect QNearFieldShareTarget::error")

	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_DisconnectError(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error")
	}
}

//export callbackQNearFieldShareTargetError
func callbackQNearFieldShareTargetError(ptr unsafe.Pointer, ptrName *C.char, error C.int) {
	defer qt.Recovering("callback QNearFieldShareTarget::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error"); signal != nil {
		signal.(func(QNearFieldShareManager__ShareError))(QNearFieldShareManager__ShareError(error))
	}

}

func (ptr *QNearFieldShareTarget) Error(error QNearFieldShareManager__ShareError) {
	defer qt.Recovering("QNearFieldShareTarget::error")

	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_Error(ptr.Pointer(), C.int(error))
	}
}

func (ptr *QNearFieldShareTarget) IsShareInProgress() bool {
	defer qt.Recovering("QNearFieldShareTarget::isShareInProgress")

	if ptr.Pointer() != nil {
		return C.QNearFieldShareTarget_IsShareInProgress(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNearFieldShareTarget) Share(message QNdefMessage_ITF) bool {
	defer qt.Recovering("QNearFieldShareTarget::share")

	if ptr.Pointer() != nil {
		return C.QNearFieldShareTarget_Share(ptr.Pointer(), PointerFromQNdefMessage(message)) != 0
	}
	return false
}

func (ptr *QNearFieldShareTarget) ShareError() QNearFieldShareManager__ShareError {
	defer qt.Recovering("QNearFieldShareTarget::shareError")

	if ptr.Pointer() != nil {
		return QNearFieldShareManager__ShareError(C.QNearFieldShareTarget_ShareError(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNearFieldShareTarget) ConnectShareFinished(f func()) {
	defer qt.Recovering("connect QNearFieldShareTarget::shareFinished")

	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_ConnectShareFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "shareFinished", f)
	}
}

func (ptr *QNearFieldShareTarget) DisconnectShareFinished() {
	defer qt.Recovering("disconnect QNearFieldShareTarget::shareFinished")

	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_DisconnectShareFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "shareFinished")
	}
}

//export callbackQNearFieldShareTargetShareFinished
func callbackQNearFieldShareTargetShareFinished(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QNearFieldShareTarget::shareFinished")

	if signal := qt.GetSignal(C.GoString(ptrName), "shareFinished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QNearFieldShareTarget) ShareFinished() {
	defer qt.Recovering("QNearFieldShareTarget::shareFinished")

	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_ShareFinished(ptr.Pointer())
	}
}

func (ptr *QNearFieldShareTarget) ShareModes() QNearFieldShareManager__ShareMode {
	defer qt.Recovering("QNearFieldShareTarget::shareModes")

	if ptr.Pointer() != nil {
		return QNearFieldShareManager__ShareMode(C.QNearFieldShareTarget_ShareModes(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNearFieldShareTarget) DestroyQNearFieldShareTarget() {
	defer qt.Recovering("QNearFieldShareTarget::~QNearFieldShareTarget")

	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_DestroyQNearFieldShareTarget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNearFieldShareTarget) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QNearFieldShareTarget::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QNearFieldShareTarget) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QNearFieldShareTarget::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQNearFieldShareTargetTimerEvent
func callbackQNearFieldShareTargetTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldShareTarget::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQNearFieldShareTargetFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QNearFieldShareTarget) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QNearFieldShareTarget::timerEvent")

	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QNearFieldShareTarget) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QNearFieldShareTarget::timerEvent")

	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QNearFieldShareTarget) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QNearFieldShareTarget::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QNearFieldShareTarget) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QNearFieldShareTarget::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQNearFieldShareTargetChildEvent
func callbackQNearFieldShareTargetChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldShareTarget::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQNearFieldShareTargetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QNearFieldShareTarget) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QNearFieldShareTarget::childEvent")

	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QNearFieldShareTarget) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QNearFieldShareTarget::childEvent")

	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QNearFieldShareTarget) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QNearFieldShareTarget::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QNearFieldShareTarget) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QNearFieldShareTarget::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQNearFieldShareTargetCustomEvent
func callbackQNearFieldShareTargetCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldShareTarget::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQNearFieldShareTargetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QNearFieldShareTarget) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QNearFieldShareTarget::customEvent")

	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QNearFieldShareTarget) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QNearFieldShareTarget::customEvent")

	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}