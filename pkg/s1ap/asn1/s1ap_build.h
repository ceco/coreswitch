void
S1SetupRequestBuild(S1AP_PDU_t *pdu, const char* enb_name);
void
S1SetupResponseBuild(S1AP_PDU_t *pdu, int num_served_gummei);
void
S1InitialUeMessageRequestBuild(S1AP_PDU_t *pdu, unsigned char *mmebuf, int mmebuf_len, long enb_ue_s1ap_id);
void
DownlinkNASTransportBuild(S1AP_PDU_t *pdu, long enb_ie_s1ap_id, unsigned char *mmebuf, int mmebuf_len);
void
UplinkNASTransportBuild(S1AP_PDU_t *pdu, long mme_ie_s1ap_id, long enb_ie_s1ap_id, unsigned char *mmebuf, int mmebuf_len);
void
InitialContextSetupRequestBuild(S1AP_PDU_t *pdu, long enb_ie_s1ap_id);
void
InitialContextSetupResponseBuild(S1AP_PDU_t *pdu, long mme_ie_s1ap_id, long enb_ie_s1ap_id,
unsigned char *tr_layer_add_buf, unsigned char *gtp_teid_buf);
void
UEContextReleaseCommandBuild(S1AP_PDU_t *pdu, long mme_ie_s1ap_id, long enb_ie_s1ap_id);
