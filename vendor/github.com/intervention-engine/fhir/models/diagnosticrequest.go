// Copyright (c) 2011-2015, HL7, Inc & The MITRE Corporation
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without modification,
// are permitted provided that the following conditions are met:
//
//     * Redistributions of source code must retain the above copyright notice, this
//       list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above copyright notice,
//       this list of conditions and the following disclaimer in the documentation
//       and/or other materials provided with the distribution.
//     * Neither the name of HL7 nor the names of its contributors may be used to
//       endorse or promote products derived from this software without specific
//       prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.
// IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT,
// INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT
// NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
// PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
// WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

type DiagnosticRequest struct {
	DomainResource        `bson:",inline"`
	Identifier            []Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Definition            []Reference       `bson:"definition,omitempty" json:"definition,omitempty"`
	BasedOn               []Reference       `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	Replaces              []Reference       `bson:"replaces,omitempty" json:"replaces,omitempty"`
	Requisition           *Identifier       `bson:"requisition,omitempty" json:"requisition,omitempty"`
	Status                string            `bson:"status,omitempty" json:"status,omitempty"`
	Stage                 *CodeableConcept  `bson:"stage,omitempty" json:"stage,omitempty"`
	Code                  *CodeableConcept  `bson:"code,omitempty" json:"code,omitempty"`
	Subject               *Reference        `bson:"subject,omitempty" json:"subject,omitempty"`
	Context               *Reference        `bson:"context,omitempty" json:"context,omitempty"`
	OccurrenceDateTime    *FHIRDateTime     `bson:"occurrenceDateTime,omitempty" json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod      *Period           `bson:"occurrencePeriod,omitempty" json:"occurrencePeriod,omitempty"`
	OccurrenceTiming      *Timing           `bson:"occurrenceTiming,omitempty" json:"occurrenceTiming,omitempty"`
	Authored              *FHIRDateTime     `bson:"authored,omitempty" json:"authored,omitempty"`
	Requester             *Reference        `bson:"requester,omitempty" json:"requester,omitempty"`
	PerformerType         *CodeableConcept  `bson:"performerType,omitempty" json:"performerType,omitempty"`
	Performer             *Reference        `bson:"performer,omitempty" json:"performer,omitempty"`
	Reason                []CodeableConcept `bson:"reason,omitempty" json:"reason,omitempty"`
	SupportingInformation []Reference       `bson:"supportingInformation,omitempty" json:"supportingInformation,omitempty"`
	Note                  []Annotation      `bson:"note,omitempty" json:"note,omitempty"`
	RelevantHistory       []Reference       `bson:"relevantHistory,omitempty" json:"relevantHistory,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *DiagnosticRequest) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "DiagnosticRequest"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to DiagnosticRequest), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *DiagnosticRequest) GetBSON() (interface{}, error) {
	x.ResourceType = "DiagnosticRequest"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "diagnosticRequest" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type diagnosticRequest DiagnosticRequest

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *DiagnosticRequest) UnmarshalJSON(data []byte) (err error) {
	x2 := diagnosticRequest{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = DiagnosticRequest(x2)
		return x.checkResourceType()
	}
	return
}

func (x *DiagnosticRequest) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "DiagnosticRequest"
	} else if x.ResourceType != "DiagnosticRequest" {
		return errors.New(fmt.Sprintf("Expected resourceType to be DiagnosticRequest, instead received %s", x.ResourceType))
	}
	return nil
}

type DiagnosticRequestPlus struct {
	DiagnosticRequest                     `bson:",inline"`
	DiagnosticRequestPlusRelatedResources `bson:",inline"`
}

type DiagnosticRequestPlusRelatedResources struct {
	IncludedPractitionerResourcesReferencedByRequester             *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByRequester,omitempty"`
	IncludedOrganizationResourcesReferencedByRequester             *[]Organization          `bson:"_includedOrganizationResourcesReferencedByRequester,omitempty"`
	IncludedDeviceResourcesReferencedByRequester                   *[]Device                `bson:"_includedDeviceResourcesReferencedByRequester,omitempty"`
	IncludedGroupResourcesReferencedBySubject                      *[]Group                 `bson:"_includedGroupResourcesReferencedBySubject,omitempty"`
	IncludedDeviceResourcesReferencedBySubject                     *[]Device                `bson:"_includedDeviceResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedBySubject                    *[]Patient               `bson:"_includedPatientResourcesReferencedBySubject,omitempty"`
	IncludedLocationResourcesReferencedBySubject                   *[]Location              `bson:"_includedLocationResourcesReferencedBySubject,omitempty"`
	IncludedEpisodeOfCareResourcesReferencedByEncounter            *[]EpisodeOfCare         `bson:"_includedEpisodeOfCareResourcesReferencedByEncounter,omitempty"`
	IncludedEncounterResourcesReferencedByEncounter                *[]Encounter             `bson:"_includedEncounterResourcesReferencedByEncounter,omitempty"`
	IncludedPatientResourcesReferencedByPatient                    *[]Patient               `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedPractitionerResourcesReferencedByFiller                *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByFiller,omitempty"`
	IncludedOrganizationResourcesReferencedByFiller                *[]Organization          `bson:"_includedOrganizationResourcesReferencedByFiller,omitempty"`
	IncludedDeviceResourcesReferencedByFiller                      *[]Device                `bson:"_includedDeviceResourcesReferencedByFiller,omitempty"`
	IncludedPatientResourcesReferencedByFiller                     *[]Patient               `bson:"_includedPatientResourcesReferencedByFiller,omitempty"`
	IncludedRelatedPersonResourcesReferencedByFiller               *[]RelatedPerson         `bson:"_includedRelatedPersonResourcesReferencedByFiller,omitempty"`
	RevIncludedReferralRequestResourcesReferencingBasedon          *[]ReferralRequest       `bson:"_revIncludedReferralRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref      *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref      *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedConsentResourcesReferencingData                     *[]Consent               `bson:"_revIncludedConsentResourcesReferencingData,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref     *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedContractResourcesReferencingTtopic                  *[]Contract              `bson:"_revIncludedContractResourcesReferencingTtopic,omitempty"`
	RevIncludedContractResourcesReferencingSubject                 *[]Contract              `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedContractResourcesReferencingTopic                   *[]Contract              `bson:"_revIncludedContractResourcesReferencingTopic,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequestreference   *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingRequestreference,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponsereference  *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingResponsereference,omitempty"`
	RevIncludedImplementationGuideResourcesReferencingResource     *[]ImplementationGuide   `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
	RevIncludedImagingStudyResourcesReferencingBasedon             *[]ImagingStudy          `bson:"_revIncludedImagingStudyResourcesReferencingBasedon,omitempty"`
	RevIncludedCommunicationResourcesReferencingBasedon            *[]Communication         `bson:"_revIncludedCommunicationResourcesReferencingBasedon,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData               *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingFocus                       *[]Task                  `bson:"_revIncludedTaskResourcesReferencingFocus,omitempty"`
	RevIncludedCarePlanResourcesReferencingActivityreference       *[]CarePlan              `bson:"_revIncludedCarePlanResourcesReferencingActivityreference,omitempty"`
	RevIncludedListResourcesReferencingItem                        *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingReplaces       *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingBasedon        *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingDefinition     *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingDefinition,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingReplaces        *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingBasedon         *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingDefinition      *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingDefinition,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                    *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedDiagnosticReportResourcesReferencingRequest         *[]DiagnosticReport      `bson:"_revIncludedDiagnosticReportResourcesReferencingRequest,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity                *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject              *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated         *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingBasedon    *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingBasedon,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject    *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequestreference *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequestreference,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingAction        *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingAction,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingPlan          *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingPlan,omitempty"`
}

func (d *DiagnosticRequestPlusRelatedResources) GetIncludedPractitionerResourceReferencedByRequester() (practitioner *Practitioner, err error) {
	if d.IncludedPractitionerResourcesReferencedByRequester == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*d.IncludedPractitionerResourcesReferencedByRequester) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*d.IncludedPractitionerResourcesReferencedByRequester))
	} else if len(*d.IncludedPractitionerResourcesReferencedByRequester) == 1 {
		practitioner = &(*d.IncludedPractitionerResourcesReferencedByRequester)[0]
	}
	return
}

func (d *DiagnosticRequestPlusRelatedResources) GetIncludedOrganizationResourceReferencedByRequester() (organization *Organization, err error) {
	if d.IncludedOrganizationResourcesReferencedByRequester == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*d.IncludedOrganizationResourcesReferencedByRequester) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*d.IncludedOrganizationResourcesReferencedByRequester))
	} else if len(*d.IncludedOrganizationResourcesReferencedByRequester) == 1 {
		organization = &(*d.IncludedOrganizationResourcesReferencedByRequester)[0]
	}
	return
}

func (d *DiagnosticRequestPlusRelatedResources) GetIncludedDeviceResourceReferencedByRequester() (device *Device, err error) {
	if d.IncludedDeviceResourcesReferencedByRequester == nil {
		err = errors.New("Included devices not requested")
	} else if len(*d.IncludedDeviceResourcesReferencedByRequester) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*d.IncludedDeviceResourcesReferencedByRequester))
	} else if len(*d.IncludedDeviceResourcesReferencedByRequester) == 1 {
		device = &(*d.IncludedDeviceResourcesReferencedByRequester)[0]
	}
	return
}

func (d *DiagnosticRequestPlusRelatedResources) GetIncludedGroupResourceReferencedBySubject() (group *Group, err error) {
	if d.IncludedGroupResourcesReferencedBySubject == nil {
		err = errors.New("Included groups not requested")
	} else if len(*d.IncludedGroupResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*d.IncludedGroupResourcesReferencedBySubject))
	} else if len(*d.IncludedGroupResourcesReferencedBySubject) == 1 {
		group = &(*d.IncludedGroupResourcesReferencedBySubject)[0]
	}
	return
}

func (d *DiagnosticRequestPlusRelatedResources) GetIncludedDeviceResourceReferencedBySubject() (device *Device, err error) {
	if d.IncludedDeviceResourcesReferencedBySubject == nil {
		err = errors.New("Included devices not requested")
	} else if len(*d.IncludedDeviceResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*d.IncludedDeviceResourcesReferencedBySubject))
	} else if len(*d.IncludedDeviceResourcesReferencedBySubject) == 1 {
		device = &(*d.IncludedDeviceResourcesReferencedBySubject)[0]
	}
	return
}

func (d *DiagnosticRequestPlusRelatedResources) GetIncludedPatientResourceReferencedBySubject() (patient *Patient, err error) {
	if d.IncludedPatientResourcesReferencedBySubject == nil {
		err = errors.New("Included patients not requested")
	} else if len(*d.IncludedPatientResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*d.IncludedPatientResourcesReferencedBySubject))
	} else if len(*d.IncludedPatientResourcesReferencedBySubject) == 1 {
		patient = &(*d.IncludedPatientResourcesReferencedBySubject)[0]
	}
	return
}

func (d *DiagnosticRequestPlusRelatedResources) GetIncludedLocationResourceReferencedBySubject() (location *Location, err error) {
	if d.IncludedLocationResourcesReferencedBySubject == nil {
		err = errors.New("Included locations not requested")
	} else if len(*d.IncludedLocationResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*d.IncludedLocationResourcesReferencedBySubject))
	} else if len(*d.IncludedLocationResourcesReferencedBySubject) == 1 {
		location = &(*d.IncludedLocationResourcesReferencedBySubject)[0]
	}
	return
}

func (d *DiagnosticRequestPlusRelatedResources) GetIncludedEpisodeOfCareResourceReferencedByEncounter() (episodeOfCare *EpisodeOfCare, err error) {
	if d.IncludedEpisodeOfCareResourcesReferencedByEncounter == nil {
		err = errors.New("Included episodeofcares not requested")
	} else if len(*d.IncludedEpisodeOfCareResourcesReferencedByEncounter) > 1 {
		err = fmt.Errorf("Expected 0 or 1 episodeOfCare, but found %d", len(*d.IncludedEpisodeOfCareResourcesReferencedByEncounter))
	} else if len(*d.IncludedEpisodeOfCareResourcesReferencedByEncounter) == 1 {
		episodeOfCare = &(*d.IncludedEpisodeOfCareResourcesReferencedByEncounter)[0]
	}
	return
}

func (d *DiagnosticRequestPlusRelatedResources) GetIncludedEncounterResourceReferencedByEncounter() (encounter *Encounter, err error) {
	if d.IncludedEncounterResourcesReferencedByEncounter == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*d.IncludedEncounterResourcesReferencedByEncounter) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*d.IncludedEncounterResourcesReferencedByEncounter))
	} else if len(*d.IncludedEncounterResourcesReferencedByEncounter) == 1 {
		encounter = &(*d.IncludedEncounterResourcesReferencedByEncounter)[0]
	}
	return
}

func (d *DiagnosticRequestPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if d.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*d.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*d.IncludedPatientResourcesReferencedByPatient))
	} else if len(*d.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*d.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (d *DiagnosticRequestPlusRelatedResources) GetIncludedPractitionerResourceReferencedByFiller() (practitioner *Practitioner, err error) {
	if d.IncludedPractitionerResourcesReferencedByFiller == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*d.IncludedPractitionerResourcesReferencedByFiller) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*d.IncludedPractitionerResourcesReferencedByFiller))
	} else if len(*d.IncludedPractitionerResourcesReferencedByFiller) == 1 {
		practitioner = &(*d.IncludedPractitionerResourcesReferencedByFiller)[0]
	}
	return
}

func (d *DiagnosticRequestPlusRelatedResources) GetIncludedOrganizationResourceReferencedByFiller() (organization *Organization, err error) {
	if d.IncludedOrganizationResourcesReferencedByFiller == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*d.IncludedOrganizationResourcesReferencedByFiller) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*d.IncludedOrganizationResourcesReferencedByFiller))
	} else if len(*d.IncludedOrganizationResourcesReferencedByFiller) == 1 {
		organization = &(*d.IncludedOrganizationResourcesReferencedByFiller)[0]
	}
	return
}

func (d *DiagnosticRequestPlusRelatedResources) GetIncludedDeviceResourceReferencedByFiller() (device *Device, err error) {
	if d.IncludedDeviceResourcesReferencedByFiller == nil {
		err = errors.New("Included devices not requested")
	} else if len(*d.IncludedDeviceResourcesReferencedByFiller) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*d.IncludedDeviceResourcesReferencedByFiller))
	} else if len(*d.IncludedDeviceResourcesReferencedByFiller) == 1 {
		device = &(*d.IncludedDeviceResourcesReferencedByFiller)[0]
	}
	return
}

func (d *DiagnosticRequestPlusRelatedResources) GetIncludedPatientResourceReferencedByFiller() (patient *Patient, err error) {
	if d.IncludedPatientResourcesReferencedByFiller == nil {
		err = errors.New("Included patients not requested")
	} else if len(*d.IncludedPatientResourcesReferencedByFiller) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*d.IncludedPatientResourcesReferencedByFiller))
	} else if len(*d.IncludedPatientResourcesReferencedByFiller) == 1 {
		patient = &(*d.IncludedPatientResourcesReferencedByFiller)[0]
	}
	return
}

func (d *DiagnosticRequestPlusRelatedResources) GetIncludedRelatedPersonResourceReferencedByFiller() (relatedPerson *RelatedPerson, err error) {
	if d.IncludedRelatedPersonResourcesReferencedByFiller == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*d.IncludedRelatedPersonResourcesReferencedByFiller) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*d.IncludedRelatedPersonResourcesReferencedByFiller))
	} else if len(*d.IncludedRelatedPersonResourcesReferencedByFiller) == 1 {
		relatedPerson = &(*d.IncludedRelatedPersonResourcesReferencedByFiller)[0]
	}
	return
}

func (d *DiagnosticRequestPlusRelatedResources) GetRevIncludedReferralRequestResourcesReferencingBasedon() (referralRequests []ReferralRequest, err error) {
	if d.RevIncludedReferralRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded referralRequests not requested")
	} else {
		referralRequests = *d.RevIncludedReferralRequestResourcesReferencingBasedon
	}
	return
}

func (d *DiagnosticRequestPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if d.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *d.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (d *DiagnosticRequestPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if d.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *d.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (d *DiagnosticRequestPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if d.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *d.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (d *DiagnosticRequestPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if d.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *d.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (d *DiagnosticRequestPlusRelatedResources) GetRevIncludedContractResourcesReferencingTtopic() (contracts []Contract, err error) {
	if d.RevIncludedContractResourcesReferencingTtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *d.RevIncludedContractResourcesReferencingTtopic
	}
	return
}

func (d *DiagnosticRequestPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if d.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *d.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (d *DiagnosticRequestPlusRelatedResources) GetRevIncludedContractResourcesReferencingTopic() (contracts []Contract, err error) {
	if d.RevIncludedContractResourcesReferencingTopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *d.RevIncludedContractResourcesReferencingTopic
	}
	return
}

func (d *DiagnosticRequestPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequestreference() (paymentNotices []PaymentNotice, err error) {
	if d.RevIncludedPaymentNoticeResourcesReferencingRequestreference == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *d.RevIncludedPaymentNoticeResourcesReferencingRequestreference
	}
	return
}

func (d *DiagnosticRequestPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponsereference() (paymentNotices []PaymentNotice, err error) {
	if d.RevIncludedPaymentNoticeResourcesReferencingResponsereference == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *d.RevIncludedPaymentNoticeResourcesReferencingResponsereference
	}
	return
}

func (d *DiagnosticRequestPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if d.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *d.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (d *DiagnosticRequestPlusRelatedResources) GetRevIncludedImagingStudyResourcesReferencingBasedon() (imagingStudies []ImagingStudy, err error) {
	if d.RevIncludedImagingStudyResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded imagingStudies not requested")
	} else {
		imagingStudies = *d.RevIncludedImagingStudyResourcesReferencingBasedon
	}
	return
}

func (d *DiagnosticRequestPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if d.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *d.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (d *DiagnosticRequestPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if d.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *d.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (d *DiagnosticRequestPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if d.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *d.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (d *DiagnosticRequestPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if d.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *d.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (d *DiagnosticRequestPlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingActivityreference() (carePlans []CarePlan, err error) {
	if d.RevIncludedCarePlanResourcesReferencingActivityreference == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *d.RevIncludedCarePlanResourcesReferencingActivityreference
	}
	return
}

func (d *DiagnosticRequestPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if d.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *d.RevIncludedListResourcesReferencingItem
	}
	return
}

func (d *DiagnosticRequestPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingReplaces() (diagnosticRequests []DiagnosticRequest, err error) {
	if d.RevIncludedDiagnosticRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *d.RevIncludedDiagnosticRequestResourcesReferencingReplaces
	}
	return
}

func (d *DiagnosticRequestPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingBasedon() (diagnosticRequests []DiagnosticRequest, err error) {
	if d.RevIncludedDiagnosticRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *d.RevIncludedDiagnosticRequestResourcesReferencingBasedon
	}
	return
}

func (d *DiagnosticRequestPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingDefinition() (diagnosticRequests []DiagnosticRequest, err error) {
	if d.RevIncludedDiagnosticRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *d.RevIncludedDiagnosticRequestResourcesReferencingDefinition
	}
	return
}

func (d *DiagnosticRequestPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingReplaces() (deviceUseRequests []DeviceUseRequest, err error) {
	if d.RevIncludedDeviceUseRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *d.RevIncludedDeviceUseRequestResourcesReferencingReplaces
	}
	return
}

func (d *DiagnosticRequestPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingBasedon() (deviceUseRequests []DeviceUseRequest, err error) {
	if d.RevIncludedDeviceUseRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *d.RevIncludedDeviceUseRequestResourcesReferencingBasedon
	}
	return
}

func (d *DiagnosticRequestPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingDefinition() (deviceUseRequests []DeviceUseRequest, err error) {
	if d.RevIncludedDeviceUseRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *d.RevIncludedDeviceUseRequestResourcesReferencingDefinition
	}
	return
}

func (d *DiagnosticRequestPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if d.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *d.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (d *DiagnosticRequestPlusRelatedResources) GetRevIncludedDiagnosticReportResourcesReferencingRequest() (diagnosticReports []DiagnosticReport, err error) {
	if d.RevIncludedDiagnosticReportResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded diagnosticReports not requested")
	} else {
		diagnosticReports = *d.RevIncludedDiagnosticReportResourcesReferencingRequest
	}
	return
}

func (d *DiagnosticRequestPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if d.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *d.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (d *DiagnosticRequestPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if d.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *d.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (d *DiagnosticRequestPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if d.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *d.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (d *DiagnosticRequestPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if d.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *d.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (d *DiagnosticRequestPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingBasedon() (questionnaireResponses []QuestionnaireResponse, err error) {
	if d.RevIncludedQuestionnaireResponseResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *d.RevIncludedQuestionnaireResponseResourcesReferencingBasedon
	}
	return
}

func (d *DiagnosticRequestPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if d.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *d.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (d *DiagnosticRequestPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequestreference() (processResponses []ProcessResponse, err error) {
	if d.RevIncludedProcessResponseResourcesReferencingRequestreference == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *d.RevIncludedProcessResponseResourcesReferencingRequestreference
	}
	return
}

func (d *DiagnosticRequestPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingAction() (clinicalImpressions []ClinicalImpression, err error) {
	if d.RevIncludedClinicalImpressionResourcesReferencingAction == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *d.RevIncludedClinicalImpressionResourcesReferencingAction
	}
	return
}

func (d *DiagnosticRequestPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingPlan() (clinicalImpressions []ClinicalImpression, err error) {
	if d.RevIncludedClinicalImpressionResourcesReferencingPlan == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *d.RevIncludedClinicalImpressionResourcesReferencingPlan
	}
	return
}

func (d *DiagnosticRequestPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if d.IncludedPractitionerResourcesReferencedByRequester != nil {
		for idx := range *d.IncludedPractitionerResourcesReferencedByRequester {
			rsc := (*d.IncludedPractitionerResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedOrganizationResourcesReferencedByRequester != nil {
		for idx := range *d.IncludedOrganizationResourcesReferencedByRequester {
			rsc := (*d.IncludedOrganizationResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedDeviceResourcesReferencedByRequester != nil {
		for idx := range *d.IncludedDeviceResourcesReferencedByRequester {
			rsc := (*d.IncludedDeviceResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *d.IncludedGroupResourcesReferencedBySubject {
			rsc := (*d.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedDeviceResourcesReferencedBySubject != nil {
		for idx := range *d.IncludedDeviceResourcesReferencedBySubject {
			rsc := (*d.IncludedDeviceResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *d.IncludedPatientResourcesReferencedBySubject {
			rsc := (*d.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedLocationResourcesReferencedBySubject != nil {
		for idx := range *d.IncludedLocationResourcesReferencedBySubject {
			rsc := (*d.IncludedLocationResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedEpisodeOfCareResourcesReferencedByEncounter != nil {
		for idx := range *d.IncludedEpisodeOfCareResourcesReferencedByEncounter {
			rsc := (*d.IncludedEpisodeOfCareResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *d.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*d.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *d.IncludedPatientResourcesReferencedByPatient {
			rsc := (*d.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedPractitionerResourcesReferencedByFiller != nil {
		for idx := range *d.IncludedPractitionerResourcesReferencedByFiller {
			rsc := (*d.IncludedPractitionerResourcesReferencedByFiller)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedOrganizationResourcesReferencedByFiller != nil {
		for idx := range *d.IncludedOrganizationResourcesReferencedByFiller {
			rsc := (*d.IncludedOrganizationResourcesReferencedByFiller)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedDeviceResourcesReferencedByFiller != nil {
		for idx := range *d.IncludedDeviceResourcesReferencedByFiller {
			rsc := (*d.IncludedDeviceResourcesReferencedByFiller)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedPatientResourcesReferencedByFiller != nil {
		for idx := range *d.IncludedPatientResourcesReferencedByFiller {
			rsc := (*d.IncludedPatientResourcesReferencedByFiller)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedRelatedPersonResourcesReferencedByFiller != nil {
		for idx := range *d.IncludedRelatedPersonResourcesReferencedByFiller {
			rsc := (*d.IncludedRelatedPersonResourcesReferencedByFiller)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (d *DiagnosticRequestPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if d.RevIncludedReferralRequestResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedReferralRequestResourcesReferencingBasedon {
			rsc := (*d.RevIncludedReferralRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *d.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*d.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *d.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*d.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *d.RevIncludedConsentResourcesReferencingData {
			rsc := (*d.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *d.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*d.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *d.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*d.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedContractResourcesReferencingSubject {
			rsc := (*d.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *d.RevIncludedContractResourcesReferencingTopic {
			rsc := (*d.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPaymentNoticeResourcesReferencingRequestreference != nil {
		for idx := range *d.RevIncludedPaymentNoticeResourcesReferencingRequestreference {
			rsc := (*d.RevIncludedPaymentNoticeResourcesReferencingRequestreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPaymentNoticeResourcesReferencingResponsereference != nil {
		for idx := range *d.RevIncludedPaymentNoticeResourcesReferencingResponsereference {
			rsc := (*d.RevIncludedPaymentNoticeResourcesReferencingResponsereference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *d.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*d.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedImagingStudyResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedImagingStudyResourcesReferencingBasedon {
			rsc := (*d.RevIncludedImagingStudyResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*d.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *d.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*d.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *d.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*d.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *d.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*d.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCarePlanResourcesReferencingActivityreference != nil {
		for idx := range *d.RevIncludedCarePlanResourcesReferencingActivityreference {
			rsc := (*d.RevIncludedCarePlanResourcesReferencingActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *d.RevIncludedListResourcesReferencingItem {
			rsc := (*d.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *d.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*d.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDiagnosticRequestResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedDiagnosticRequestResourcesReferencingBasedon {
			rsc := (*d.RevIncludedDiagnosticRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDiagnosticRequestResourcesReferencingDefinition != nil {
		for idx := range *d.RevIncludedDiagnosticRequestResourcesReferencingDefinition {
			rsc := (*d.RevIncludedDiagnosticRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *d.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*d.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDeviceUseRequestResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedDeviceUseRequestResourcesReferencingBasedon {
			rsc := (*d.RevIncludedDeviceUseRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDeviceUseRequestResourcesReferencingDefinition != nil {
		for idx := range *d.RevIncludedDeviceUseRequestResourcesReferencingDefinition {
			rsc := (*d.RevIncludedDeviceUseRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*d.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDiagnosticReportResourcesReferencingRequest != nil {
		for idx := range *d.RevIncludedDiagnosticReportResourcesReferencingRequest {
			rsc := (*d.RevIncludedDiagnosticReportResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *d.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*d.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*d.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *d.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*d.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *d.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*d.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedQuestionnaireResponseResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedQuestionnaireResponseResourcesReferencingBasedon {
			rsc := (*d.RevIncludedQuestionnaireResponseResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*d.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedProcessResponseResourcesReferencingRequestreference != nil {
		for idx := range *d.RevIncludedProcessResponseResourcesReferencingRequestreference {
			rsc := (*d.RevIncludedProcessResponseResourcesReferencingRequestreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedClinicalImpressionResourcesReferencingAction != nil {
		for idx := range *d.RevIncludedClinicalImpressionResourcesReferencingAction {
			rsc := (*d.RevIncludedClinicalImpressionResourcesReferencingAction)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedClinicalImpressionResourcesReferencingPlan != nil {
		for idx := range *d.RevIncludedClinicalImpressionResourcesReferencingPlan {
			rsc := (*d.RevIncludedClinicalImpressionResourcesReferencingPlan)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (d *DiagnosticRequestPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if d.IncludedPractitionerResourcesReferencedByRequester != nil {
		for idx := range *d.IncludedPractitionerResourcesReferencedByRequester {
			rsc := (*d.IncludedPractitionerResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedOrganizationResourcesReferencedByRequester != nil {
		for idx := range *d.IncludedOrganizationResourcesReferencedByRequester {
			rsc := (*d.IncludedOrganizationResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedDeviceResourcesReferencedByRequester != nil {
		for idx := range *d.IncludedDeviceResourcesReferencedByRequester {
			rsc := (*d.IncludedDeviceResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *d.IncludedGroupResourcesReferencedBySubject {
			rsc := (*d.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedDeviceResourcesReferencedBySubject != nil {
		for idx := range *d.IncludedDeviceResourcesReferencedBySubject {
			rsc := (*d.IncludedDeviceResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *d.IncludedPatientResourcesReferencedBySubject {
			rsc := (*d.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedLocationResourcesReferencedBySubject != nil {
		for idx := range *d.IncludedLocationResourcesReferencedBySubject {
			rsc := (*d.IncludedLocationResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedEpisodeOfCareResourcesReferencedByEncounter != nil {
		for idx := range *d.IncludedEpisodeOfCareResourcesReferencedByEncounter {
			rsc := (*d.IncludedEpisodeOfCareResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *d.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*d.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *d.IncludedPatientResourcesReferencedByPatient {
			rsc := (*d.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedPractitionerResourcesReferencedByFiller != nil {
		for idx := range *d.IncludedPractitionerResourcesReferencedByFiller {
			rsc := (*d.IncludedPractitionerResourcesReferencedByFiller)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedOrganizationResourcesReferencedByFiller != nil {
		for idx := range *d.IncludedOrganizationResourcesReferencedByFiller {
			rsc := (*d.IncludedOrganizationResourcesReferencedByFiller)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedDeviceResourcesReferencedByFiller != nil {
		for idx := range *d.IncludedDeviceResourcesReferencedByFiller {
			rsc := (*d.IncludedDeviceResourcesReferencedByFiller)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedPatientResourcesReferencedByFiller != nil {
		for idx := range *d.IncludedPatientResourcesReferencedByFiller {
			rsc := (*d.IncludedPatientResourcesReferencedByFiller)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedRelatedPersonResourcesReferencedByFiller != nil {
		for idx := range *d.IncludedRelatedPersonResourcesReferencedByFiller {
			rsc := (*d.IncludedRelatedPersonResourcesReferencedByFiller)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedReferralRequestResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedReferralRequestResourcesReferencingBasedon {
			rsc := (*d.RevIncludedReferralRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *d.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*d.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *d.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*d.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *d.RevIncludedConsentResourcesReferencingData {
			rsc := (*d.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *d.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*d.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *d.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*d.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedContractResourcesReferencingSubject {
			rsc := (*d.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *d.RevIncludedContractResourcesReferencingTopic {
			rsc := (*d.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPaymentNoticeResourcesReferencingRequestreference != nil {
		for idx := range *d.RevIncludedPaymentNoticeResourcesReferencingRequestreference {
			rsc := (*d.RevIncludedPaymentNoticeResourcesReferencingRequestreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPaymentNoticeResourcesReferencingResponsereference != nil {
		for idx := range *d.RevIncludedPaymentNoticeResourcesReferencingResponsereference {
			rsc := (*d.RevIncludedPaymentNoticeResourcesReferencingResponsereference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *d.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*d.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedImagingStudyResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedImagingStudyResourcesReferencingBasedon {
			rsc := (*d.RevIncludedImagingStudyResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*d.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *d.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*d.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *d.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*d.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *d.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*d.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCarePlanResourcesReferencingActivityreference != nil {
		for idx := range *d.RevIncludedCarePlanResourcesReferencingActivityreference {
			rsc := (*d.RevIncludedCarePlanResourcesReferencingActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *d.RevIncludedListResourcesReferencingItem {
			rsc := (*d.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *d.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*d.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDiagnosticRequestResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedDiagnosticRequestResourcesReferencingBasedon {
			rsc := (*d.RevIncludedDiagnosticRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDiagnosticRequestResourcesReferencingDefinition != nil {
		for idx := range *d.RevIncludedDiagnosticRequestResourcesReferencingDefinition {
			rsc := (*d.RevIncludedDiagnosticRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *d.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*d.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDeviceUseRequestResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedDeviceUseRequestResourcesReferencingBasedon {
			rsc := (*d.RevIncludedDeviceUseRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDeviceUseRequestResourcesReferencingDefinition != nil {
		for idx := range *d.RevIncludedDeviceUseRequestResourcesReferencingDefinition {
			rsc := (*d.RevIncludedDeviceUseRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*d.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDiagnosticReportResourcesReferencingRequest != nil {
		for idx := range *d.RevIncludedDiagnosticReportResourcesReferencingRequest {
			rsc := (*d.RevIncludedDiagnosticReportResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *d.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*d.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*d.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *d.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*d.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *d.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*d.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedQuestionnaireResponseResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedQuestionnaireResponseResourcesReferencingBasedon {
			rsc := (*d.RevIncludedQuestionnaireResponseResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*d.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedProcessResponseResourcesReferencingRequestreference != nil {
		for idx := range *d.RevIncludedProcessResponseResourcesReferencingRequestreference {
			rsc := (*d.RevIncludedProcessResponseResourcesReferencingRequestreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedClinicalImpressionResourcesReferencingAction != nil {
		for idx := range *d.RevIncludedClinicalImpressionResourcesReferencingAction {
			rsc := (*d.RevIncludedClinicalImpressionResourcesReferencingAction)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedClinicalImpressionResourcesReferencingPlan != nil {
		for idx := range *d.RevIncludedClinicalImpressionResourcesReferencingPlan {
			rsc := (*d.RevIncludedClinicalImpressionResourcesReferencingPlan)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
