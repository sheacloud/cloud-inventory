import { useFetch } from '../hooks';
import { Container, Col, Row } from 'react-bootstrap';
import { NavLink, Outlet } from 'react-router-dom';

export default function Aws() {
    const url = `http://localhost:3000/v1/metadata/aws/`;
    const { status, data, error } = useFetch(url);
    return (
        <Container fluid>
            <Row className="flex-nowrap">
                <Col md="auto">
                    <h1>AWS</h1>
                    {status === 'fetching' && <p>Loading services...</p>}
                    {status === 'error' && <p>Error: {error}</p>}
                    {status === 'fetched' && (
                        <nav style={{ borderRight: 'solid 1px', padding: '1rem' }}>
                            {data.services.map(service => (
                                <NavLink key={service} to={`/aws/${service}`} style={({ isActive }) => {
                                    return {
                                        display: 'block',
                                        margin: '1rem 0',
                                        color: isActive ? 'red' : '',
                                    };
                                }}>
                                    {service}
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