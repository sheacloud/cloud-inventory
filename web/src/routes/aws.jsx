import { useFetch } from '../hooks';
import { Container, Col, Row } from 'react-bootstrap';
import { NavLink, Outlet } from 'react-router-dom';

export default function Aws() {
    const url = `http://localhost:3000/v1/metadata/aws/`;
    const { status, data, error } = useFetch(url);
    return (
        <Container fluid style={{ display: 'flex', flexDirection: 'column', flex: '1' }}>
            <Row className="flex-nowrap" style={{ display: 'flex', flex: '1' }}>
                <Col md="auto" className="sticky-sidebar" style={{ borderRight: 'solid 1px' }}>
                    <div className="sticky-top">
                        <h1>AWS</h1>
                        {status === 'fetching' && <p>Loading services...</p>}
                        {status === 'error' && <p>Error: {error}</p>}
                        {status === 'fetched' && (
                            <nav style={{ paddingRight: '1rem', display: 'flex', flexDirection: 'column', flex: '1' }}>
                                {data.services.map(service => (
                                    <NavLink key={service} to={`/aws/${service}`} style={({ isActive }) => {
                                        return {
                                            display: 'block',
                                            margin: '2px 0',
                                            color: isActive ? 'green' : '',
                                        };
                                    }}>
                                        {service}
                                    </NavLink>
                                ))
                                }
                            </nav>
                        )}
                    </div>
                </Col>
                <Col style={{ display: 'flex', flexDirection: 'column', flex: '1' }}>
                    <Outlet />
                </Col>
            </Row>
        </Container>
    )
}