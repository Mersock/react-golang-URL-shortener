import React, { useState, useEffect } from 'react'
import { Form, FormGroup, Input, Table } from 'reactstrap'
import { useDispatch } from 'react-redux'
import { GET_URL } from '../../actions/types'
import urlShortens from '../../apis/urlShortens'
import { getUrl } from '../../actions'


const TableList = ({ tableList }) => {
    const dispatch = useDispatch()
    const [term, setTerm] = useState('')
    const [debouncedTerm, setDebouncedTerm] = useState(term)

    const handleTbody = () => {
        const list = tableList.map((list, index) => {
            const { counter, expires, originalUrl, shortUrl, urlCode } = list
            return (
                <tr key={index}>
                    <th scope="row">{index + 1}</th>
                    <td>{urlCode}</td>
                    <td><a href={shortUrl} target="_blank" rel="noreferrer" >{shortUrl} </a></td>
                    <td><a href={originalUrl} target="_blank" rel="noreferrer" >{originalUrl}</a></td>
                    <td>{expires}</td>
                    <td>{counter}</td>
                </tr>
            )
        })
        return list
    }

    useEffect(() => {
        const timeoutId = setTimeout(() => {
            setDebouncedTerm(term)
        }, 1000);

        return () => {
            clearTimeout(timeoutId)
        }

    }, [term])

    useEffect(() => {
        const fetchSearch = async () => {
            try {
                const res = await urlShortens.get('/api/urlShorten', { params: { urlCode: debouncedTerm } })
                dispatch({ type: GET_URL, payload: res.data });
            } catch (error) {
                console.error("getUrl", error);
            }
        }

        if (debouncedTerm) {
            fetchSearch()
        } else {
            dispatch(getUrl())
        }

    }, [debouncedTerm, dispatch])

    return (
        <div style={{ marginTop: '20px' }}>
            <Form>
                <FormGroup>
                    <Input type="text" name="urlCode" id="urlCode" placeholder="Url Code" onChange={(e) => setTerm(e.target.value)} />
                </FormGroup>
            </Form>
            <Table hover striped bordered borderless>
                <thead>
                    <tr>
                        <th>#</th>
                        <th>Url Code</th>
                        <th>Short Url</th>
                        <th>Original Url</th>
                        <th>expires</th>
                        <th>Counter</th>
                    </tr>
                </thead>
                <tbody>
                    {handleTbody()}
                </tbody>
            </Table>
        </div>
    )
}

export default TableList
