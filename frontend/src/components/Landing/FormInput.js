import React, { useEffect, useState } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { Button, Form, FormGroup, Label, Input, Alert, FormFeedback } from 'reactstrap';
import { createUrl } from '../../actions';
import _ from 'lodash'

const FormInput = (props) => {
    const dispatch = useDispatch()
    const [inputInValid, setinputInValid] = useState(false)
    const [urlInput, seturlInput] = useState('')
    const [shortUrl, setshortUrl] = useState('')
    const urlShorten = useSelector(state => state.urlShortens)

    const handleSubmit = (e) => {
        e.preventDefault()
        const url = isUrl(urlInput)
        if (!url) {
            setinputInValid(true)
        } else {
            dispatch(createUrl({
                OriginalUrl: _.trim(urlInput)
            }))
            setinputInValid(false)
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

    useEffect(() => {
        const { shortUrl } = urlShorten
        if (!_.isUndefined(shortUrl)) {
            setshortUrl(shortUrl)
        }
    }, [urlShorten])

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
                    <Input onChange={(e) => seturlInput(e.target.value)} invalid={inputInValid} type="text" name="originalURL" id="originalURL" placeholder="originalURL" />
                    <FormFeedback>Oh noes! that Url is invalid</FormFeedback>
                </FormGroup>
                <FormGroup>
                    {
                        shortUrl !== '' ? (
                            <div>
                                <Label for="exampleShortURL"><h2>Short Url</h2></Label>
                                <Alert color="success">
                                    <a href={shortUrl} target="_blank" without rel="noreferrer" > {shortUrl}</a>
                                </Alert>
                            </div>
                        ) : null
                    }

                </FormGroup>
                <Button color="primary" size="lg" block>Submit</Button>
            </Form>
        </div>
    )
}

export default FormInput
