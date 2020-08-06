package qqmeeting

import "encoding/json"

func (meeting Meeting) handleResponse(respBody []byte, descriptor *MeetingRequestDescriptor) (MeetingResponse, error) {
	var response MeetingResponse
	var err error
	switch *descriptor {
	case RequestDescriptorMeetingCreate:
		var resp MeetingCreateResponse
		err = json.Unmarshal(respBody, &resp)
		response = resp
	case RequestDescriptorUserCreate, RequestDescriptorUserUpdate, RequestDescriptorUserDelete:
		// 返回的body为空
		return nil, nil
	case RequestDescriptorUserDetailQuery:
		var resp UserDetailQueryResponse
		err = json.Unmarshal(respBody, &resp)
		response = resp
	case RequestDescriptorUserList:
		var resp UserListResponse
		err = json.Unmarshal(respBody, &resp)
		response = resp
	}

	if err != nil {
		return nil, err
	} else {
		return response, nil
	}
}
