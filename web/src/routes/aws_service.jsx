import { useFetch } from '../hooks';
import { Container, Col, Row } from 'react-bootstrap';
import { NavLink, Outlet, useParams } from 'react-router-dom';

export default function AwsService() {
    let params = useParams();
    const url = `http://localhost:3000/v1/metadata/aws/${params.service}`;
    const { status, data, error } = useFetch(url);
    return (
        <Container fluid>
            <Row className="flex-nowrap">
                <Col md="auto">
                    <h1>{params.service}</h1>
                    {status === 'fetching' && <p>Loading resources...</p>}
                    {status === 'error' && <p>Error: {error}</p>}
                    {status === 'fetched' && (
                        <nav style={{ borderRight: 'solid 1px', padding: '1rem' }}>
                            {data.resources.map(resource => (
                                <NavLink key={resource} to={`/aws/${params.service}/${resource}`} style={({ isActive }) => {
                                    return {
                                        display: 'block',
                                        margin: '1rem 0',
                                        color: isActive ? 'red' : '',
                                    };
                                }}>
                                    {resource}
                                </NavLink>
                            ))
                            }
                        </nav>
                    )}
                </Col>
                <Col>
                    <Outlet />
                </Col>
            </Row>
        </Container>
    )
}