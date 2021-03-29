import React, { useEffect } from 'react'
import { Col, Row } from 'reactstrap'
import TableList from './tableList'
import { getUrl } from '../../actions'
import { useDispatch, useSelector } from 'react-redux'

const Index = (props) => {
    const dispatch = useDispatch()
    const tableList = useSelector(state => state.urlShortens.listUrl)

    useEffect(() => {
        dispatch(getUrl())
        //eslint-disable-next-line react-hooks/exhaustive-deps
    }, [])

    return (
        <Row>
            <Col>
                <TableList tableList={tableList} />
            </Col>
        </Row>
    )
}

export default Index
