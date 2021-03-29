import React, { useState } from 'react'
import { Button, Form, FormGroup, Label, Input, Alert, FormFeedback } from 'reactstrap';

const FormInput = (props) => {
    const [inputValid, setinputValid] = useState(false)
    const [urlInput, seturlInput] = useState('')

    const handleSubmit = (e) => {
        e.preventDefault()
        const url = isUrl(urlInput)
        if (!url) {
            setinputValid(true)
        } else {
            setinputValid(false)
        }

    }

    const isUrl = (val) => {
        const expression = /(https?:\/\/(?:www\.|(?!www))[a-zA-Z0-9][a-zA-Z0-9-]+[a-zA-Z0-9]\.[^\s]{2,}|www\.[a-zA-Z0-9][a-zA-Z0-9-]+[a-zA-Z0-9]\.[^\s]{2,}|https?:\/\/(?:www\.|(?!www))[a-zA-Z0-9]+\.[^\s]{2,}|www\.[a-zA-Z0-9]+\.[^\s]{2,})/gi;
        const regex = new RegExp(expression);
        if (val.match(regex)) {
            return true
        }
        return false
    }

    return (
        <div>
            <Form onSubmit={handleSubmit}>
                <FormGroup>
                    <Label for="originalURL">
                        <h2>Original Url</h2>
                        <div>
                            <span>Ex: <a href="https://www.youtube.com">https://www.youtube.com</a></span>
                        </div>
                    </Label>
                    <Input onChange={(e) => seturlInput(e.target.value)} invalid={inputValid} type="text" name="originalURL" id="originalURL" placeholder="originalURL" />
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
