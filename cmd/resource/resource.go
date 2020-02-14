package resource

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
)


// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
    // Add your code here:
    // * Make API calls (use req.Session)
    // * Mutate the model
    // * Check/set any callback context (req.CallbackContext / response.CallbackContext)

    /*
        // Construct a new handler.ProgressEvent and return it
        response := handler.ProgressEvent{
            OperationStatus: handler.Success,
            Message: "Create complete",
            ResourceModel: currentModel,
        }

        return response, nil
    */

		// Example sending a request using StartAutomationExecutionRequest.
		req := client.StartAutomationExecutionRequest(currentModel)
		resp, err := req.Send(context.TODO())
		if err == nil {
		  fmt.Println(resp)
		}
		response := handler.ProgressEvent{
				OperationStatus: handler.Success,
				Message: "Create complete",
				ResourceModel: currentModel,
		}

		return response, nil

    // Not implemented, return an empty handler.ProgressEvent
    // and an error
    return handler.ProgressEvent{}, errors.New("Not implemented: Create")
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
    // Add your code here:
    // * Make API calls (use req.Session)
    // * Mutate the model
    // * Check/set any callback context (req.CallbackContext / response.CallbackContext)

    /*
        // Construct a new handler.ProgressEvent and return it
        response := handler.ProgressEvent{
            OperationStatus: handler.Success,
            Message: "Read complete",
            ResourceModel: currentModel,
        }

        return response, nil
    */

    // Not implemented, return an empty handler.ProgressEvent
    // and an error
    return handler.ProgressEvent{}, errors.New("Not implemented: Read")
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
    // Add your code here:
    // * Make API calls (use req.Session)
    // * Mutate the model
    // * Check/set any callback context (req.CallbackContext / response.CallbackContext)

    /*
        // Construct a new handler.ProgressEvent and return it
        response := handler.ProgressEvent{
            OperationStatus: handler.Success,
            Message: "Update complete",
            ResourceModel: currentModel,
        }

        return response, nil
    */

    // Not implemented, return an empty handler.ProgressEvent
    // and an error
    return handler.ProgressEvent{}, errors.New("Not implemented: Update")
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
    // Add your code here:
    // * Make API calls (use req.Session)
    // * Mutate the model
    // * Check/set any callback context (req.CallbackContext / response.CallbackContext)

    /*
        // Construct a new handler.ProgressEvent and return it
        response := handler.ProgressEvent{
            OperationStatus: handler.Success,
            Message: "Delete complete",
            ResourceModel: currentModel,
        }

        return response, nil
    */

    // Not implemented, return an empty handler.ProgressEvent
    // and an error
    return handler.ProgressEvent{}, errors.New("Not implemented: Delete")
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
    // Add your code here:
    // * Make API calls (use req.Session)
    // * Mutate the model
    // * Check/set any callback context (req.CallbackContext / response.CallbackContext)

    /*
        // Construct a new handler.ProgressEvent and return it
        response := handler.ProgressEvent{
            OperationStatus: handler.Success,
            Message: "List complete",
            ResourceModel: currentModel,
        }

        return response, nil
    */

    // Not implemented, return an empty handler.ProgressEvent
    // and an error
    return handler.ProgressEvent{}, errors.New("Not implemented: List")
}
