package pipedrive

import (
	"context"
	"fmt"
	"net/http"
)

// ActivitiesService handles activities related
// methods of the Pipedrive API.
//
// Pipedrive API dcos: https://developers.pipedrive.com/docs/api/v1/#!/Activities
type ActivitiesService service

// Participants represents a Pipedrive participant.
type Participants struct {
	PersonID    int  `json:"person_id"`
	PrimaryFlag bool `json:"primary_flag"`
}

// Activity represents a Pipedrive activity.
type Activity struct {
	Id                       int         `json:"id"`
	Type                     string      `json:"type"`
	Duration                 string      `json:"duration"`
	Subject                  string      `json:"subject"`
	CompanyID                int         `json:"company_id"`
	UserID                   int         `json:"user_id"`
	Done                     bool        `json:"done"`
	DueDate                  string      `json:"due_date"`
	DueTime                  string      `json:"due_time"`
	AddTime                  string      `json:"add_time"`
	MarkedAsDoneTime         string      `json:"marked_as_done_time"`
	OrgID                    int         `json:"org_id"`
	PersonID                 int         `json:"person_id"`
	DealID                   int         `json:"deal_id"`
	ActiveFlag               bool        `json:"active_flag"`
	UpdateTime               string      `json:"update_time"`
	ConferenceMeetingClient  interface{} `json:"conference_meeting_client"`
	ConferenceMeetingURL     interface{} `json:"conference_meeting_url"`
	ConferenceMeetingID      int         `json:"conference_meeting_id"`
	BusyFlag                 bool        `json:"busy_flag"`
	PublicDescription        string      `json:"public_description"`
	Location                 string      `json:"location"`
	UpdateUserID             int         `json:"update_user_id"`
	SourceTimezone           string      `json:"source_timezone"`
	LeadID                   int         `json:"lead_id"`
	LocationSubpremise       interface{} `json:"location_subpremise"`
	LocationStreetNumber     int         `json:"location_street_number"`
	LocationRoute            string      `json:"location_route"`
	LocationSublocality      string      `json:"location_sublocality"`
	LocationLocality         string      `json:"location_locality"`
	LocationAdminAreaLevel1  string      `json:"location_admin_area_level_1"`
	LocationAdminAreaLevel2  string      `json:"location_admin_area_level_2"`
	LocationCountry          string      `json:"location_country"`
	LocationPostalCode       string      `json:"location_postal_code"`
	LocationFormattedAddress string      `json:"location_formatted_address"`
	ProjectID                int         `json:"project_id"`
}

func (a Activity) String() string {
	return Stringify(a)
}

// ActivityResponse represents single activity response.
type ActivityResponse struct {
	Success bool     `json:"success"`
	Data    Activity `json:"data"`
}

// ActivitiesReponse represents multiple activities response.
type ActivitiesReponse struct {
	Success        bool           `json:"success"`
	Data           []Activity     `json:"data"`
	AdditionalData AdditionalData `json:"additional_data,omitempty"`
}

// List returns total count users
func (s *ActivitiesService) Summary(ctx context.Context) (*Summary, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/activities/summary", nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *Summary

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// List returns all activities assigned to a particular user
//
// https://developers.pipedrive.com/docs/api/v1/#!/Activities/get_activities
func (s *ActivitiesService) List(ctx context.Context, opts PaginationParameters) (*ActivitiesReponse, *Response, error) {
	var (
		err error
		req *http.Request
	)

	switch {
	case opts.Limit > 0 || len(opts.Cursor) != 0:
		req, err = s.client.NewRequest(http.MethodGet, "/activities/collection", &opts, nil)
	default:
		req, err = s.client.NewRequest(http.MethodGet, "/activities/collection", nil, nil)
	}

	var record *ActivitiesReponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// GetByID returns details of a specific activity.
//
// https://developers.pipedrive.com/docs/api/v1/#!/Activities/get_activities
func (s *ActivitiesService) GetByID(ctx context.Context, id int) (*ActivitiesReponse, *Response, error) {
	uri := fmt.Sprintf("/activities/%v", id)
	req, err := s.client.NewRequest(http.MethodGet, uri, nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *ActivitiesReponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Create an activity.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Activities/post_activities
func (s *ActivitiesService) Create(ctx context.Context, opt *ActivitiesCreateOptions) (*ActivityResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodPost, "/activities", nil, opt)

	if err != nil {
		return nil, nil, err
	}

	var record *ActivityResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// ActivitiesCreateOptions specifices the optional parameters to the
// ActivitiesService.Update method.
type ActivitiesCreateOptions struct {
	Subject      string      `json:"subject,omitempty"`
	Done         uint8       `json:"done,omitempty"`
	Type         string      `json:"type,omitempty"`
	DueDate      string      `json:"due_date,omitempty"`
	DueTime      string      `json:"due_time,omitempty"`
	Duration     string      `json:"duration,omitempty"`
	UserID       uint        `json:"user_id,omitempty"`
	DealID       uint        `json:"user_id,omitempty"`
	PersonID     uint        `json:"person_id,omitempty"`
	Participants interface{} `json:"participants,omitempty"`
	OrgID        uint        `json:"org_id,omitempty"`
}

// Update an activity
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Activities/put_activities_id
func (s *ActivitiesService) Update(ctx context.Context, id int, opt *ActivitiesCreateOptions) (*ActivityResponse, *Response, error) {
	uri := fmt.Sprintf("/activities/%v", id)
	req, err := s.client.NewRequest(http.MethodPut, uri, opt, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *ActivityResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// DeleteMultiple activities in bulk.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Activities/delete_activities
func (s *ActivitiesService) DeleteMultiple(ctx context.Context, ids []int) (*Response, error) {
	req, err := s.client.NewRequest(http.MethodDelete, "/activities", &DeleteMultipleOptions{
		Ids: arrayToString(ids, ","),
	}, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// Delete an activity.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Activities/delete_activities_id
func (s *ActivitiesService) Delete(ctx context.Context, id int) (*Response, error) {
	uri := fmt.Sprintf("/activities/%v", id)
	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}
