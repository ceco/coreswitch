/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "S1AP-IEs"
 * 	found in "r14.4.0/36413-e40.asn"
 * 	`asn1c -pdu=all -fcompound-names -findirect-choice -fno-include-deps -no-gen-example`
 */

#ifndef	_ImmediateMDT_H_
#define	_ImmediateMDT_H_


#include <asn_application.h>

/* Including external dependencies */
#include "MeasurementsToActivate.h"
#include "M1ReportingTrigger.h"
#include <constr_SEQUENCE.h>

#ifdef __cplusplus
extern "C" {
#endif

/* Forward declarations */
struct M1ThresholdEventA2;
struct M1PeriodicReporting;
struct ProtocolExtensionContainer;

/* ImmediateMDT */
typedef struct ImmediateMDT {
	MeasurementsToActivate_t	 measurementsToActivate;
	M1ReportingTrigger_t	 m1reportingTrigger;
	struct M1ThresholdEventA2	*m1thresholdeventA2;	/* OPTIONAL */
	struct M1PeriodicReporting	*m1periodicReporting;	/* OPTIONAL */
	struct ProtocolExtensionContainer	*iE_Extensions;	/* OPTIONAL */
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} ImmediateMDT_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_ImmediateMDT;
extern asn_SEQUENCE_specifics_t asn_SPC_ImmediateMDT_specs_1;
extern asn_TYPE_member_t asn_MBR_ImmediateMDT_1[5];

#ifdef __cplusplus
}
#endif

#endif	/* _ImmediateMDT_H_ */
#include <asn_internal.h>