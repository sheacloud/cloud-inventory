import { useState, useEffect } from 'react';
import { useFetch } from '../hooks';
import { Container, Col, Row, Accordion, Form, Tabs } from 'react-bootstrap';
import { useParams } from 'react-router-dom';
import SyntaxHighlighter from 'react-syntax-highlighter';
import { monokaiSublime } from 'react-syntax-highlighter/dist/esm/styles/hljs';

function GetResourceName(resource, displayNames) {
    var displayName = ""
    for (var i = 0; i < displayNames.length; i++) {
        displayName += resource[displayNames[i]]
        if (i < displayNames.length - 1) {
            displayName += " "
        }
    }
    return displayName
}

export default function AwsResource() {

    return (
        <Container fluid>
            <Tabs defaultActiveKey="inventory">
                <Tabs.Tab eventKey="inventory" title="Inventory">
                    <AwsResourceInventory />
                </Tabs.Tab>
                <Tabs.Tab eventKey="diff" title="Diff">
                    <AwsResourceDiff />
                </Tabs.Tab>
            </Tabs>
        </Container>
    )
}

function AwsResourceInventory() {
    let params = useParams();
    console.log("aws resource running", params.resource)
    const currentDate = new Date();

    const [metadataUrl, setMetadataUrl] = useState();
    const [inventoryUrl, setInventoryUrl] = useState();

    const metadataFetch = useFetch(metadataUrl);
    const inventoryFetch = useFetch(inventoryUrl);

    const [reportDate, setReportDate] = useState(currentDate.toISOString().split('T')[0]);
    const [reportTime, setReportTime] = useState('');


    useEffect(() => {
        console.log("params updated", params.resource)
        const baseMetadataUrl = `http://localhost:3000/v1/metadata/aws/${params.service}/${params.resource}`;
        const baseInventoryUrl = `http://localhost:3000/v1/inventory/aws/${params.service}/${params.resource}`;
        setMetadataUrl(baseMetadataUrl);
        setInventoryUrl(baseInventoryUrl);
    }, [params])

    useEffect(() => {
        var newInventoryUrl = `http://localhost:3000/v1/inventory/aws/${params.service}/${params.resource}`
        var newMetadataUrl = `http://localhost:3000/v1/metadata/aws/${params.service}/${params.resource}`
        if (reportDate !== '') {
            newInventoryUrl += `?report_date=${reportDate}`;
            newMetadataUrl += `?report_date=${reportDate}`;
        }
        if (reportTime !== '') {
            newInventoryUrl += `&time_selection=at&time_selection_reference=${reportTime}`;
        }
        setInventoryUrl(newInventoryUrl);
        setMetadataUrl(newMetadataUrl);
    }, [reportDate, reportTime])

    return (
        <Container fluid>
            <h1>{params.resource}</h1>
            <input type="date" defaultValue={currentDate.toISOString().split('T')[0]} onChange={(e) => setReportDate(e.target.value)} />
            {metadataFetch.status === 'fetching' && <p>Loading metadata...</p>}
            {metadataFetch.status === 'error' && <p>Error: {metadataFetch.error}</p>}
            {metadataFetch.status === 'fetched' && (
                <Container fluid>
                    <Form>
                        <Form.Group>
                            <Form.Label>Inventory Timestamp</Form.Label>
                            <Form.Select onChange={(e) => { setReportTime(e.target.value) }} defaultValue={"default"} name="module">
                                <option value={"default"} disabled>Select a specific timestamp</option>
                                {metadataFetch.data.datetimes.map((datetime, index) => {
                                    let date = new Date(datetime);
                                    return (
                                        <option key={datetime} value={datetime}>{date.toISOString()}</option>
                                    )
                                })}
                            </Form.Select>
                        </Form.Group>
                    </Form>
                    <br />
                    <h2>Inventory</h2>
                    <Container fluid style={{ overflow: 'auto', maxHeight: '60vh' }}>
                        {inventoryFetch.status === 'fetching' && <p>Loading inventory...</p>}
                        {inventoryFetch.status === 'error' && <p>Error: {inventoryFetch.error}</p>}
                        {inventoryFetch.status === 'fetched' && (
                            <Accordion>
                                {inventoryFetch.data[params.resource] && inventoryFetch.data[params.resource].map((resource, index) => (
                                    <Accordion.Item eventKey={resource[metadataFetch.data.id_field]}>
                                        <Accordion.Header>{GetResourceName(resource, metadataFetch.data.display_fields)}</Accordion.Header>
                                        <Accordion.Body>
                                            <SyntaxHighlighter language="json" style={monokaiSublime}>
                                                {JSON.stringify(resource, null, 2)}
                                            </SyntaxHighlighter>
                                        </Accordion.Body>
                                    </Accordion.Item>
                                ))}
                            </Accordion>
                        )}
                    </Container>
                </Container>
            )}
        </Container>
    )
}



function AwsResourceDiff() {
    let params = useParams();
    console.log("aws diff running", params.resource)
    const currentDate = new Date();

    const [startMetadataUrl, setStartMetadataUrl] = useState();
    const [endMetadataUrl, setEndMetadataUrl] = useState();
    const [diffUrl, setDiffUrl] = useState();

    const startMetadataFetch = useFetch(startMetadataUrl);
    const endMetadataFetch = useFetch(endMetadataUrl);
    const diffFetch = useFetch(diffUrl);

    const [startReportDate, setStartReportDate] = useState(currentDate.toISOString().split('T')[0]);
    const [endReportDate, setEndReportDate] = useState(currentDate.toISOString().split('T')[0]);

    const [startReportTime, setStartReportTime] = useState('');
    const [endReportTime, setEndReportTime] = useState('');

    useEffect(() => {
        console.log("params updated", params.resource)
        const baseMetadataUrl = `http://localhost:3000/v1/metadata/aws/${params.service}/${params.resource}`;
        setStartMetadataUrl(baseMetadataUrl);
        setEndMetadataUrl(baseMetadataUrl);
    }, [params])

    useEffect(() => {
        var newDiffUrl = `http://localhost:3000/v1/diff/aws/${params.service}/${params.resource}`
        var newStartMetadataUrl = `http://localhost:3000/v1/metadata/aws/${params.service}/${params.resource}`
        var newEndMetadataUrl = `http://localhost:3000/v1/metadata/aws/${params.service}/${params.resource}`

        if (startReportDate !== '') {
            newStartMetadataUrl += `?report_date=${startReportDate}`;
            newDiffUrl += `?start_report_date=${startReportDate}`;
        }
        if (endReportDate !== '') {
            newEndMetadataUrl += `?report_date=${endReportDate}`;
            newDiffUrl += `&end_report_date=${endReportDate}`;
        }
        if (startReportTime !== '') {
            newDiffUrl += `&start_time_selection=at&start_time_selection_reference=${startReportTime}`;
        }
        if (endReportTime !== '') {
            newDiffUrl += `&end_time_selection=at&end_time_selection_reference=${endReportTime}`;
        }

        if (startReportDate !== '' && endReportDate !== '' && startReportTime !== '' && endReportTime !== '') {
            setDiffUrl(newDiffUrl);
        }

        setStartMetadataUrl(newStartMetadataUrl);
        setEndMetadataUrl(newEndMetadataUrl);
    }, [startReportDate, endReportDate, startReportTime, endReportTime])

    let processedDiff = {};
    if (diffFetch.status === 'fetched') {
        processedDiff = processDiff(diffFetch.data)
    }
    return (
        <Container fluid>
            <h1>{params.resource}</h1>
            <Row>
                <Col md={4}>
                    <h2>Diff Start Time</h2>
                    <input type="date" defaultValue={currentDate.toISOString().split('T')[0]} onChange={(e) => setStartReportDate(e.target.value)} />
                    {startMetadataFetch.status === 'fetching' && <p>Loading metadata...</p>}
                    {startMetadataFetch.status === 'error' && <p>Error: {startMetadataFetch.error}</p>}
                    {startMetadataFetch.status === 'fetched' && (
                        <Container fluid>
                            <Form>
                                <Form.Group>
                                    <Form.Label>Start Timestamp</Form.Label>
                                    <Form.Select onChange={(e) => { setStartReportTime(e.target.value) }} defaultValue={"default"} name="module">
                                        <option value={"default"} disabled>Select a specific timestamp</option>
                                        {startMetadataFetch.data.datetimes.map((datetime, index) => {
                                            let date = new Date(datetime);
                                            return (
                                                <option key={datetime} value={datetime}>{date.toISOString()}</option>
                                            )
                                        })}
                                    </Form.Select>
                                </Form.Group>
                            </Form>
                        </Container>
                    )}
                </Col>
                <Col />
                <Col md={4}>
                    <h2>Diff End Time</h2>
                    <input type="date" defaultValue={currentDate.toISOString().split('T')[0]} onChange={(e) => setEndReportDate(e.target.value)} />
                    {endMetadataFetch.status === 'fetching' && <p>Loading metadata...</p>}
                    {endMetadataFetch.status === 'error' && <p>Error: {endMetadataFetch.error}</p>}
                    {endMetadataFetch.status === 'fetched' && (
                        <Container fluid>
                            <Form>
                                <Form.Group>
                                    <Form.Label>End Timestamp</Form.Label>
                                    <Form.Select onChange={(e) => { setEndReportTime(e.target.value) }} defaultValue={"default"} name="module">
                                        <option value={"default"} disabled>Select a specific timestamp</option>
                                        {endMetadataFetch.data.datetimes.map((datetime, index) => {
                                            let date = new Date(datetime);
                                            return (
                                                <option key={datetime} value={datetime}>{date.toISOString()}</option>
                                            )
                                        })}
                                    </Form.Select>
                                </Form.Group>
                            </Form>
                        </Container>
                    )}
                </Col>
            </Row>
            {diffFetch.status === 'idle' && <p>Please select a start and end time.</p>}
            {diffFetch.status === 'fetching' && <p>Loading diff...</p>}
            {diffFetch.status === 'error' && <p>Error: {diffFetch.error}</p>}
            {diffFetch.status === 'fetched' && (
                <Container fluid>
                    <br />
                    <p style={{ 'white-space': 'pre-wrap' }}>{JSON.stringify(processedDiff, null, 2)}</p>
                </Container>
            )}
        </Container>
    )
}

function processDiff(diff) {
    const diffs = {
        created: [],
        updated: [],
        deleted: [],
    };
    diff.forEach(function (item) {
        if (item.type === 'create') {
            diffs.created.push({ path: item.path, value: item.to });
        } else if (item.type === 'update') {
            diffs.updated.push({ path: item.path, from: item.from, to: item.to });
        } else if (item.type === 'delete') {
            diffs.deleted.push({ path: item.path, value: item.from });
        }
    });
    return diffs;
}