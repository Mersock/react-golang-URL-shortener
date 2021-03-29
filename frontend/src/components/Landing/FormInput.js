import React, { useState } from 'react'
import { Button, Form, FormGroup, Label, Input, Alert, FormFeedback } from 'reactstrap';

const FormInput = (props) => {
    const [inputValid, setinputValid] = useState(false)
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
                    <Input invalid={inputValid} type="text" name="originalURL" id="originalURL" placeholder="originalURL" />
                    <FormFeedback>Oh noes! that Url is invalid</FormFeedback>
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
