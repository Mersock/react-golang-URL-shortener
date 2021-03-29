import React from 'react'
import { Form, FormGroup, Input, Table } from 'reactstrap'

const TableList = ({ tableList }) => {

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

    return (
        <div>
            <Form>
                <FormGroup>
                    <Input type="text" name="urlCode" id="urlCode" placeholder="Url Code" />
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
