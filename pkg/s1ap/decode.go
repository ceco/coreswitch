package s1ap

// #cgo CFLAGS: -I./asn1
// #cgo LDFLAGS: -L/usr/local/lib -ls1ap
// #include "S1AP-PDU.h"
// #include "InitiatingMessage.h"
// #include "ProtocolIE-Field.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	"reflect"
	"unsafe"
)

var S1AP_PDU2StringMap = map[C.S1AP_PDU_PR]string{
	C.S1AP_PDU_PR_NOTHING:             "Nothing",
	C.S1AP_PDU_PR_initiatingMessage:   "InitiatingMessage",
	C.S1AP_PDU_PR_successfulOutcome:   "SuccessfulOutcome",
	C.S1AP_PDU_PR_unsuccessfulOutcome: "UnsuccessfulOutcome",
}

func S1AP_PDU2String(k C.S1AP_PDU_PR) string {
	if v, ok := S1AP_PDU2StringMap[k]; ok {
		return v
	} else {
		return "Unknown"
	}
}

var S1AP_Initiating2StringMap = map[C.InitiatingMessage__value_PR]string{
	C.InitiatingMessage__value_PR_NOTHING:                              "NOTHING",
	C.InitiatingMessage__value_PR_HandoverRequired:                     "HandoverRequired",
	C.InitiatingMessage__value_PR_HandoverRequest:                      "HandoverRequest",
	C.InitiatingMessage__value_PR_PathSwitchRequest:                    "PathSwitchRequest",
	C.InitiatingMessage__value_PR_E_RABSetupRequest:                    "E_RABSetupRequest",
	C.InitiatingMessage__value_PR_E_RABModifyRequest:                   "E_RABModifyRequest",
	C.InitiatingMessage__value_PR_E_RABReleaseCommand:                  "E_RABReleaseCommand",
	C.InitiatingMessage__value_PR_InitialContextSetupRequest:           "InitialContextSetupRequest",
	C.InitiatingMessage__value_PR_HandoverCancel:                       "HandoverCancel",
	C.InitiatingMessage__value_PR_KillRequest:                          "KillRequest",
	C.InitiatingMessage__value_PR_Reset:                                "Reset",
	C.InitiatingMessage__value_PR_S1SetupRequest:                       "S1SetupRequest",
	C.InitiatingMessage__value_PR_UEContextModificationRequest:         "UEContextModificationRequest",
	C.InitiatingMessage__value_PR_UEContextReleaseCommand:              "UEContextReleaseCommand",
	C.InitiatingMessage__value_PR_ENBConfigurationUpdate:               "ENBConfigurationUpdate",
	C.InitiatingMessage__value_PR_MMEConfigurationUpdate:               "MMEConfigurationUpdate",
	C.InitiatingMessage__value_PR_WriteReplaceWarningRequest:           "WriteReplaceWarningRequest",
	C.InitiatingMessage__value_PR_UERadioCapabilityMatchRequest:        "UERadioCapabilityMatchRequest",
	C.InitiatingMessage__value_PR_E_RABModificationIndication:          "E_RABModificationIndication",
	C.InitiatingMessage__value_PR_UEContextModificationIndication:      "UEContextModificationIndication",
	C.InitiatingMessage__value_PR_UEContextSuspendRequest:              "UEContextSuspendRequest",
	C.InitiatingMessage__value_PR_UEContextResumeRequest:               "UEContextResumeRequest",
	C.InitiatingMessage__value_PR_HandoverNotify:                       "HandoverNotify",
	C.InitiatingMessage__value_PR_E_RABReleaseIndication:               "E_RABReleaseIndication",
	C.InitiatingMessage__value_PR_Paging:                               "Paging",
	C.InitiatingMessage__value_PR_DownlinkNASTransport:                 "DownlinkNASTransport",
	C.InitiatingMessage__value_PR_InitialUEMessage:                     "InitialUEMessage",
	C.InitiatingMessage__value_PR_UplinkNASTransport:                   "UplinkNASTransport",
	C.InitiatingMessage__value_PR_ErrorIndication:                      "ErrorIndication",
	C.InitiatingMessage__value_PR_NASNonDeliveryIndication:             "NASNonDeliveryIndication",
	C.InitiatingMessage__value_PR_UEContextReleaseRequest:              "UEContextReleaseRequest",
	C.InitiatingMessage__value_PR_DownlinkS1cdma2000tunnelling:         "DownlinkS1cdma2000tunnelling",
	C.InitiatingMessage__value_PR_UplinkS1cdma2000tunnelling:           "UplinkS1cdma2000tunnelling",
	C.InitiatingMessage__value_PR_UECapabilityInfoIndication:           "UECapabilityInfoIndication",
	C.InitiatingMessage__value_PR_ENBStatusTransfer:                    "ENBStatusTransfer",
	C.InitiatingMessage__value_PR_MMEStatusTransfer:                    "MMEStatusTransfer",
	C.InitiatingMessage__value_PR_DeactivateTrace:                      "DeactivateTrace",
	C.InitiatingMessage__value_PR_TraceStart:                           "TraceStart",
	C.InitiatingMessage__value_PR_TraceFailureIndication:               "TraceFailureIndication",
	C.InitiatingMessage__value_PR_CellTrafficTrace:                     "CellTrafficTrace",
	C.InitiatingMessage__value_PR_LocationReportingControl:             "LocationReportingControl",
	C.InitiatingMessage__value_PR_LocationReportingFailureIndication:   "LocationReportingFailureIndication",
	C.InitiatingMessage__value_PR_LocationReport:                       "LocationReport",
	C.InitiatingMessage__value_PR_OverloadStart:                        "OverloadStart",
	C.InitiatingMessage__value_PR_OverloadStop:                         "OverloadStop",
	C.InitiatingMessage__value_PR_ENBDirectInformationTransfer:         "ENBDirectInformationTransfer",
	C.InitiatingMessage__value_PR_MMEDirectInformationTransfer:         "MMEDirectInformationTransfer",
	C.InitiatingMessage__value_PR_ENBConfigurationTransfer:             "ENBConfigurationTransfer",
	C.InitiatingMessage__value_PR_MMEConfigurationTransfer:             "MMEConfigurationTransfer",
	C.InitiatingMessage__value_PR_PrivateMessage:                       "PrivateMessage",
	C.InitiatingMessage__value_PR_DownlinkUEAssociatedLPPaTransport:    "DownlinkUEAssociatedLPPaTransport",
	C.InitiatingMessage__value_PR_UplinkUEAssociatedLPPaTransport:      "UplinkUEAssociatedLPPaTransport",
	C.InitiatingMessage__value_PR_DownlinkNonUEAssociatedLPPaTransport: "DownlinkNonUEAssociatedLPPaTransport",
	C.InitiatingMessage__value_PR_UplinkNonUEAssociatedLPPaTransport:   "UplinkNonUEAssociatedLPPaTransport",
	C.InitiatingMessage__value_PR_PWSRestartIndication:                 "PWSRestartIndication",
	C.InitiatingMessage__value_PR_RerouteNASRequest:                    "RerouteNASRequest",
	C.InitiatingMessage__value_PR_PWSFailureIndication:                 "PWSFailureIndication",
	C.InitiatingMessage__value_PR_ConnectionEstablishmentIndication:    "ConnectionEstablishmentIndication",
	C.InitiatingMessage__value_PR_NASDeliveryIndication:                "NASDeliveryIndication",
	C.InitiatingMessage__value_PR_RetrieveUEInformation:                "RetrieveUEInformation",
	C.InitiatingMessage__value_PR_UEInformationTransfer:                "UEInformationTransfer",
	C.InitiatingMessage__value_PR_ENBCPRelocationIndication:            "ENBCPRelocationIndication",
	C.InitiatingMessage__value_PR_MMECPRelocationIndication:            "MMECPRelocationIndication",
}

func S1AP_Initiating2String(k C.InitiatingMessage__value_PR) string {
	if v, ok := S1AP_Initiating2StringMap[k]; ok {
		return v
	} else {
		return "Unknown"
	}
}

func InitialUEMessageHandle(packet unsafe.Pointer) (int32, error) {
	pdu := (*C.S1AP_PDU_t)(packet)
	msg := *(**C.InitiatingMessage_t)(unsafe.Pointer(&pdu.choice))
	val := (*C.InitialUEMessage_t)(unsafe.Pointer(&msg.value.choice))

	var ies []*C.UplinkNASTransport_IEs_t
	slice := (*reflect.SliceHeader)((unsafe.Pointer(&ies)))
	slice.Cap = (int)(val.protocolIEs.list.count)
	slice.Len = (int)(val.protocolIEs.list.count)
	slice.Data = uintptr(unsafe.Pointer(val.protocolIEs.list.array))

	var enb_ie_s1ap_id int32

	for _, ie := range ies {
		switch ie.id {
		case C.ProtocolIE_ID_id_eNB_UE_S1AP_ID:
			enb_ie_s1ap_id_c := (*C.ENB_UE_S1AP_ID_t)(unsafe.Pointer(&ie.value.choice))
			enb_ie_s1ap_id = (int32)(*enb_ie_s1ap_id_c)
		case C.ProtocolIE_ID_id_NAS_PDU:
			//NAS_PDU = &ie->value.choice.NAS_PDU;
		case C.ProtocolIE_ID_id_TAI:
			//TAI = &ie->value.choice.TAI;
		case C.ProtocolIE_ID_id_EUTRAN_CGI:
			//EUTRAN_CGI = &ie->value.choice.EUTRAN_CGI;
		case C.ProtocolIE_ID_id_S_TMSI:
			//S_TMSI = &ie->value.choice.S_TMSI;
		default:
		}
	}
	return enb_ie_s1ap_id, nil
}

func InitialContextSetupRequestHandle(packet unsafe.Pointer) []byte {
	tmsi := []byte{}
	pdu := (*C.S1AP_PDU_t)(packet)
	msg := *(**C.InitiatingMessage_t)(unsafe.Pointer(&pdu.choice))
	val := (*C.InitialContextSetupRequest_t)(unsafe.Pointer(&msg.value.choice))

	var ies []*C.InitialContextSetupRequestIEs_t
	slice := (*reflect.SliceHeader)((unsafe.Pointer(&ies)))
	slice.Cap = (int)(val.protocolIEs.list.count)
	slice.Len = (int)(val.protocolIEs.list.count)
	slice.Data = uintptr(unsafe.Pointer(val.protocolIEs.list.array))

	for _, ie := range ies {
		switch ie.id {
		case C.ProtocolIE_ID_id_E_RABToBeSetupListCtxtSUReq:
			rab_setup := (*C.E_RABToBeSetupListCtxtSUReq_t)(unsafe.Pointer(&ie.value.choice))

			var items []*C.E_RABToBeSetupItemCtxtSUReqIEs_t
			slice := (*reflect.SliceHeader)((unsafe.Pointer(&items)))
			slice.Cap = (int)(rab_setup.list.count)
			slice.Len = (int)(rab_setup.list.count)
			slice.Data = uintptr(unsafe.Pointer(rab_setup.list.array))

			for _, item := range items {
				su_req := (*C.E_RABToBeSetupItemCtxtSUReq_t)(unsafe.Pointer(&item.value.choice))

				var nas_pdu_buf []byte
				slice := (*reflect.SliceHeader)((unsafe.Pointer(&nas_pdu_buf)))
				slice.Cap = (int)(su_req.nAS_PDU.size)
				slice.Len = (int)(su_req.nAS_PDU.size)
				slice.Data = uintptr(unsafe.Pointer(su_req.nAS_PDU.buf))

				nas_pdu_buf = nas_pdu_buf[9:]                // GPRS timer
				nas_pdu_buf = nas_pdu_buf[nas_pdu_buf[1]+1:] // TAI list
				nas_pdu_buf = nas_pdu_buf[nas_pdu_buf[2]+3:] // ESM message container

				if nas_pdu_buf[0] == 0x50 { // EPS mobility identity
					tmsi = nas_pdu_buf[2+nas_pdu_buf[1]-4 : 2+nas_pdu_buf[1]] // T-MSI
				}
			}
		}
	}

	return tmsi
}

func NAS_PDU_Handle() {
}

func ResponseHandle(packet unsafe.Pointer) int32 {
	pdu := (*C.S1AP_PDU_t)(packet)
	msg := *(**C.InitiatingMessage_t)(unsafe.Pointer(&pdu.choice))
	var eNB_UE_S1AP_ID int32 = 0

	switch msg.value.present {
	default:
		val := (*C.DownlinkNASTransport_t)(unsafe.Pointer(&msg.value.choice))

		var ies []*C.DownlinkNASTransport_IEs_t
		slice := (*reflect.SliceHeader)((unsafe.Pointer(&ies)))
		slice.Cap = (int)(val.protocolIEs.list.count)
		slice.Len = (int)(val.protocolIEs.list.count)
		slice.Data = uintptr(unsafe.Pointer(val.protocolIEs.list.array))

	FOR_LOOP:
		for _, ie := range ies {
			switch ie.id {
			case C.ProtocolIE_ID_id_eNB_UE_S1AP_ID:
				eNB_UE_S1AP_ID_c := (*C.ENB_UE_S1AP_ID_t)(unsafe.Pointer(&ie.value.choice))
				eNB_UE_S1AP_ID = (int32)(*eNB_UE_S1AP_ID_c)
				break FOR_LOOP
			}
		}

	case C.InitiatingMessage__value_PR_UEContextReleaseCommand:
		val := (*C.UEContextReleaseCommand_t)(unsafe.Pointer(&msg.value.choice))

		var items []*C.UEContextReleaseCommand_IEs_t
		slice := (*reflect.SliceHeader)((unsafe.Pointer(&items)))
		slice.Cap = (int)(val.protocolIEs.list.count)
		slice.Len = (int)(val.protocolIEs.list.count)
		slice.Data = uintptr(unsafe.Pointer(val.protocolIEs.list.array))
	FOR_LOOP_2:
		for _, item := range items {
			switch item.id {
			case C.ProtocolIE_ID_id_UE_S1AP_IDs:
				UE_S1AP_ID_c := (*C.UE_S1AP_IDs_t)(unsafe.Pointer(&item.value.choice))
				uE_S1AP_ID_pair := (*C.UE_S1AP_ID_pair_t)(unsafe.Pointer(&UE_S1AP_ID_c.choice))
				eNB_UE_S1AP_ID = (int32)((*uE_S1AP_ID_pair).eNB_UE_S1AP_ID)
				break FOR_LOOP_2
			}
		}

	}

	return eNB_UE_S1AP_ID
}

func DownlinkNASTransportReject(packet unsafe.Pointer) (byte, byte) {
	pdu := (*C.S1AP_PDU_t)(packet)
	msg := *(**C.InitiatingMessage_t)(unsafe.Pointer(&pdu.choice))
	val := (*C.DownlinkNASTransport_t)(unsafe.Pointer(&msg.value.choice))

	var ies []*C.DownlinkNASTransport_IEs_t
	slice := (*reflect.SliceHeader)((unsafe.Pointer(&ies)))
	slice.Cap = (int)(val.protocolIEs.list.count)
	slice.Len = (int)(val.protocolIEs.list.count)
	slice.Data = uintptr(unsafe.Pointer(val.protocolIEs.list.array))

	for _, ie := range ies {
		switch ie.id {
		case C.ProtocolIE_ID_id_NAS_PDU:
			nas_pdu := (*C.NAS_PDU_t)(unsafe.Pointer(&ie.value.choice))

			var nas_pdu_buf []byte
			slice := (*reflect.SliceHeader)((unsafe.Pointer(&nas_pdu_buf)))
			slice.Cap = (int)(nas_pdu.size)
			slice.Len = (int)(nas_pdu.size)
			slice.Data = uintptr(unsafe.Pointer(nas_pdu.buf))

			var securityHeaderType byte
			//var protocolDisc byte
			securityHeaderType = (nas_pdu_buf[0] & 0xf0) >> 4
			//protocolDisc = (nas_pdu_buf[0] & 0x0f)
			nas_pdu_buf = nas_pdu_buf[1:]
			//fmt.Printf("securityHeaderType: %d\n", securityHeaderType)
			//fmt.Printf("protocolDisc: %02x\n", protocolDisc)
			//fmt.Printf("Message type: %02x\n", nas_pdu_buf[6])

			if securityHeaderType == 0x00 {
				return nas_pdu_buf[0], nas_pdu_buf[1]
			}

			return nas_pdu_buf[6], nas_pdu_buf[7]
		}
	}

	return 0, 0
}

func DownlinkNASTransportHandle(packet unsafe.Pointer) (int32, []byte, uint64, uint16) {
	pdu := (*C.S1AP_PDU_t)(packet)
	msg := *(**C.InitiatingMessage_t)(unsafe.Pointer(&pdu.choice))
	val := (*C.DownlinkNASTransport_t)(unsafe.Pointer(&msg.value.choice))
	var mme_ie_s1ap_id int32 = 0

	var ies []*C.DownlinkNASTransport_IEs_t
	slice := (*reflect.SliceHeader)((unsafe.Pointer(&ies)))
	slice.Cap = (int)(val.protocolIEs.list.count)
	slice.Len = (int)(val.protocolIEs.list.count)
	slice.Data = uintptr(unsafe.Pointer(val.protocolIEs.list.array))

	for _, ie := range ies {
		switch ie.id {
		case C.ProtocolIE_ID_id_MME_UE_S1AP_ID:
			//fmt.Println("IE: MME_UE_S1AP_ID")
			mme_ie_s1ap_id_c := (*C.MME_UE_S1AP_ID_t)(unsafe.Pointer(&ie.value.choice))
			mme_ie_s1ap_id = (int32)(*mme_ie_s1ap_id_c)
		case C.ProtocolIE_ID_id_NAS_PDU:

			nas_pdu := (*C.NAS_PDU_t)(unsafe.Pointer(&ie.value.choice))

			var nas_pdu_buf []byte
			slice := (*reflect.SliceHeader)((unsafe.Pointer(&nas_pdu_buf)))
			slice.Cap = (int)(nas_pdu.size)
			slice.Len = (int)(nas_pdu.size)
			slice.Data = uintptr(unsafe.Pointer(nas_pdu.buf))

			var securityHeaderType byte
			//var protocolDisc byte

			securityHeaderType = (nas_pdu_buf[0] & 0xf0) >> 4
			//protocolDisc = (nas_pdu_buf[0] & 0x0f)
			nas_pdu_buf = nas_pdu_buf[1:]
			// typ := nas_pdu_buf[0]

			var result_code byte = 0

			if securityHeaderType == 0x00 {
				result_code = nas_pdu_buf[0]
			} else {
				result_code = nas_pdu_buf[6]
			}

			if result_code == NAS_EPS_MOBILITY_MANAGEMENT_MESSAGE_TYPE_ATTACH_REJECT || len(nas_pdu_buf) < 27 {
				return 0, []byte{}, 0, 0
			}

			/*
				fmt.Printf("securityHeaderType: %02x\n", securityHeaderType)
				fmt.Printf("protocolDisc: %02x\n", protocolDisc)
				fmt.Printf("messageType: %02x\n", typ)
				fmt.Printf("Leftover: %02x\n", nas_pdu_buf[1])
				fmt.Printf("Rand: %v\n", nas_pdu_buf[2:18])
				fmt.Printf("Len: %v\n", nas_pdu_buf[18])
				fmt.Printf("Autn: %v\n", nas_pdu_buf[19:])
				fmt.Printf("SQN: %v\n", nas_pdu_buf[19:25])
				fmt.Printf("AMF: %v\n", nas_pdu_buf[25:27])
			*/
			sqn := []byte{0, 0}
			sqn = append(sqn, nas_pdu_buf[19:25]...)

			return mme_ie_s1ap_id, nas_pdu_buf[2:18], binary.LittleEndian.Uint64(sqn), binary.LittleEndian.Uint16(nas_pdu_buf[25:27])
		}
	}

	return 0, []byte{}, 0, 0
}

func UplinkNASTransportHandle(packet unsafe.Pointer) (int32, int, error) {
	pdu := (*C.S1AP_PDU_t)(packet)
	msg := *(**C.InitiatingMessage_t)(unsafe.Pointer(&pdu.choice))
	val := (*C.UplinkNASTransport_t)(unsafe.Pointer(&msg.value.choice))

	var ies []*C.UplinkNASTransport_IEs_t
	slice := (*reflect.SliceHeader)((unsafe.Pointer(&ies)))
	slice.Cap = (int)(val.protocolIEs.list.count)
	slice.Len = (int)(val.protocolIEs.list.count)
	slice.Data = uintptr(unsafe.Pointer(val.protocolIEs.list.array))

	var enb_ie_s1ap_id int32
	var eps_mmm_type int

	for _, ie := range ies {
		switch ie.id {
		case C.ProtocolIE_ID_id_eNB_UE_S1AP_ID:
			//fmt.Println("IE: eNB_UE_S1AP_ID")
			enb_ie_s1ap_id_c := (*C.ENB_UE_S1AP_ID_t)(unsafe.Pointer(&ie.value.choice))
			enb_ie_s1ap_id = (int32)(*enb_ie_s1ap_id_c)
		case C.ProtocolIE_ID_id_NAS_PDU:
			//fmt.Println("IE: NAS_PDU")
			// OCTET_STRING_T
			// typedef struct OCTET_STRING {
			// uint8_t *buf;   /* Buffer with consecutive OCTET_STRING bits */
			// size_t size;    /* Size of the buffer */
			//
			// asn_struct_ctx_t _asn_ctx;      /* Parsing across buffer boundaries */
			// } OCTET_STRING_t;
			nas_pdu := (*C.NAS_PDU_t)(unsafe.Pointer(&ie.value.choice))

			var nas_pdu_buf []byte
			slice := (*reflect.SliceHeader)((unsafe.Pointer(&nas_pdu_buf)))
			slice.Cap = (int)(nas_pdu.size)
			slice.Len = (int)(nas_pdu.size)
			slice.Data = uintptr(unsafe.Pointer(nas_pdu.buf))
			/*
				fmt.Println("NAS_PDU_LEN", len(nas_pdu_buf))
				for _, val := range nas_pdu_buf {
					fmt.Printf("%02x ", val)
				}
				fmt.Printf("\n")
			*/
			var securityHeaderType byte
			var protocolDisc byte
			for len(nas_pdu_buf) > 0 {
				securityHeaderType = (nas_pdu_buf[0] & 0xf0) >> 4
				protocolDisc = (nas_pdu_buf[0] & 0x0f)
				nas_pdu_buf = nas_pdu_buf[1:]
				//				fmt.Printf("securityHeaderType: %02x\n", securityHeaderType)
				//				fmt.Printf("protocolDisc: %02x\n", protocolDisc)

				if protocolDisc != 7 {
					return 0, 0, fmt.Errorf("protocol discrimiter is not EPS MMM")
				}

				switch securityHeaderType {
				case 0:
					//fmt.Printf("Type %02x\n", nas_pdu_buf[0])
					typ := nas_pdu_buf[0]
					nas_pdu_buf = nas_pdu_buf[1:]
					switch typ {
					case 0x53:
						eps_mmm_type = NAS_EPS_AUTH_RESPONSE
						if len(nas_pdu_buf) > 0 {
							len := nas_pdu_buf[0]
							nas_pdu_buf = nas_pdu_buf[len+1:]
						}
					case 0x5e:
						eps_mmm_type = NAS_EPS_SECURITY_MODE_COMPLETE
					default:
						eps_mmm_type = 0
					}
				case 4:
					nas_pdu_buf = nas_pdu_buf[5:]
				default:
					return 0, 0, fmt.Errorf("security header type is not known %d", securityHeaderType)
				}
			}
		case C.ProtocolIE_ID_id_TAI:
			//fmt.Println("IE: TAI")
			//TAI = &ie->value.choice.TAI;
		case C.ProtocolIE_ID_id_EUTRAN_CGI:
			//fmt.Println("IE: EUTRAN_CGI")
			//EUTRAN_CGI = &ie->value.choice.EUTRAN_CGI;
		case C.ProtocolIE_ID_id_S_TMSI:
			//fmt.Println("IE: S_TMSI")
			//S_TMSI = &ie->value.choice.S_TMSI;
		default:
		}
	}
	return enb_ie_s1ap_id, eps_mmm_type, nil
}

func Decode(buf []byte) (unsafe.Pointer, int, error) {
	packet := C.calloc(C.sizeof_struct_S1AP_PDU, 1)
	var opt_codec *C.asn_codec_ctx_t = nil

	ret := C.aper_decode(
		opt_codec,
		&C.asn_DEF_S1AP_PDU,
		(*unsafe.Pointer)(&packet),
		(unsafe.Pointer)(&buf[0]),
		(C.size_t)(len(buf)),
		0,
		0)

	if ret.code != C.RC_OK {
		C.free(packet)
		return nil, 0, fmt.Errorf("aper_decode failed: %d", ret)
	}

	pdu := (*C.S1AP_PDU_t)(packet)
	//fmt.Println("PDU type:", S1AP_PDU2String(pdu.present))

	typ := 0
	switch pdu.present {
	case C.S1AP_PDU_PR_NOTHING:
	case C.S1AP_PDU_PR_initiatingMessage:
		msg := *(**C.InitiatingMessage_t)(unsafe.Pointer(&pdu.choice))
		//fmt.Println("Message type:", S1AP_Initiating2String(msg.value.present))
		switch msg.value.present {
		case C.InitiatingMessage__value_PR_S1SetupRequest:
			typ = S1_SETUP_REQUEST
		case C.InitiatingMessage__value_PR_InitialUEMessage:
			typ = INITIAL_UE_MESSAGE
		case C.InitiatingMessage__value_PR_InitialContextSetupRequest:
			typ = INITIAL_CONTEXT_SETUP_REQUEST
		case C.InitiatingMessage__value_PR_UplinkNASTransport:
			typ = UPLINK_NAS_TRANSPORT
		case C.InitiatingMessage__value_PR_DownlinkNASTransport:
			typ = DOWNLINK_NAS_TRANSPORT
		default:
		}
	case C.S1AP_PDU_PR_successfulOutcome:
	case C.S1AP_PDU_PR_unsuccessfulOutcome:
	default:
	}
	return packet, typ, nil
}

func XerPrint(message unsafe.Pointer) {
	C.xer_fprint(C.stdout, &C.asn_DEF_S1AP_PDU, message)
}

func Free(packet unsafe.Pointer) {
	C.free(packet)
}
