// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/CreateUserRequest
type CreateUserInput struct {
	_ struct{} `type:"structure"`

	// The path for the user name. For more information about paths, see IAM Identifiers
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html)
	// in the IAM User Guide.
	//
	// This parameter is optional. If it is not included, it defaults to a slash
	// (/).
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of either a forward slash (/) by itself
	// or a string that must begin and end with forward slashes. In addition, it
	// can contain any ASCII character from the ! (\u0021) through the DEL character
	// (\u007F), including most punctuation characters, digits, and upper and lowercased
	// letters.
	Path *string `min:"1" type:"string"`

	// The ARN of the policy that is used to set the permissions boundary for the
	// user.
	PermissionsBoundary *string `min:"20" type:"string"`

	// A list of tags that you want to attach to the newly created user. Each tag
	// consists of a key name and an associated value. For more information about
	// tagging, see Tagging IAM Identities (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html)
	// in the IAM User Guide.
	//
	// If any one of the tags is invalid or if you exceed the allowed number of
	// tags per user, then the entire request fails and the user is not created.
	Tags []Tag `type:"list"`

	// The name of the user to create.
	//
	// IAM user, group, role, and policy names must be unique within the account.
	// Names are not distinguished by case. For example, you cannot create resources
	// named both "MyResource" and "myresource".
	//
	// UserName is a required field
	UserName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateUserInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateUserInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateUserInput"}
	if s.Path != nil && len(*s.Path) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Path", 1))
	}
	if s.PermissionsBoundary != nil && len(*s.PermissionsBoundary) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("PermissionsBoundary", 20))
	}

	if s.UserName == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserName"))
	}
	if s.UserName != nil && len(*s.UserName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserName", 1))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the response to a successful CreateUser request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/CreateUserResponse
type CreateUserOutput struct {
	_ struct{} `type:"structure"`

	// A structure with details about the new IAM user.
	User *User `type:"structure"`
}

// String returns the string representation
func (s CreateUserOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateUser = "CreateUser"

// CreateUserRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Creates a new IAM user for your AWS account.
//
// For information about limitations on the number of IAM users you can create,
// see Limitations on IAM Entities (https://docs.aws.amazon.com/IAM/latest/UserGuide/LimitationsOnEntities.html)
// in the IAM User Guide.
//
//    // Example sending a request using CreateUserRequest.
//    req := client.CreateUserRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/CreateUser
func (c *Client) CreateUserRequest(input *CreateUserInput) CreateUserRequest {
	op := &aws.Operation{
		Name:       opCreateUser,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateUserInput{}
	}

	req := c.newRequest(op, input, &CreateUserOutput{})
	return CreateUserRequest{Request: req, Input: input, Copy: c.CreateUserRequest}
}

// CreateUserRequest is the request type for the
// CreateUser API operation.
type CreateUserRequest struct {
	*aws.Request
	Input *CreateUserInput
	Copy  func(*CreateUserInput) CreateUserRequest
}

// Send marshals and sends the CreateUser API request.
func (r CreateUserRequest) Send(ctx context.Context) (*CreateUserResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateUserResponse{
		CreateUserOutput: r.Request.Data.(*CreateUserOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateUserResponse is the response type for the
// CreateUser API operation.
type CreateUserResponse struct {
	*CreateUserOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateUser request.
func (r *CreateUserResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
