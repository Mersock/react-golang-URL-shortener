import React from 'react'
import { Button, Form, FormGroup, Label, Input, Alert } from 'reactstrap';

const FormInput = (props) => {
    return (
        <div>
            <Form>
                <FormGroup>
                    <Label for="originalURL">
                        <h2>Original Url</h2>
                        <div>
                            <span>Ex: <a href="https://www.youtube.com">https://www.youtube.com</a></span>
                        </div>
                    </Label>
                    <Input type="text" name="originalURL" id="originalURL" placeholder="originalURL" />
                </FormGroup>
                <FormGroup>
                    <Label for="exampleShortURL"><h2>Short Url</h2></Label>
                    <Alert color="success">
                        <a href="https://www.youtube.com"> https://www.youtube.com</a>
                    </Alert>
                </FormGroup>
                <Button color="primary" size="lg" block>Submit</Button>
            </Form>
        </div>
    )
}

export default FormInput
