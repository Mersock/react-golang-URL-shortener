import React from 'react'
import { Col, Row } from 'reactstrap'
import TableList from './tableList'
import { useSelector } from 'react-redux'

const Index = (props) => {
    const tableList = useSelector(state => state.urlShortens.listUrl)

    return (
        <Row>
            <Col>
                <TableList tableList={tableList} />
            </Col>
        </Row>
    )
}

export default Index
