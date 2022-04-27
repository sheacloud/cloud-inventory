import { useFetch } from '../hooks';
import { Container, Col, Row } from 'react-bootstrap';
import { NavLink, Outlet, useParams } from 'react-router-dom';

export default function AwsService() {
    let params = useParams();
    const url = `http://localhost:3000/v1/metadata/aws/${params.service}`;
    const { status, data, error } = useFetch(url);
    return (
        <Container fluid style={{ display: 'flex', flexDirection: 'column', flex: '1' }}>
            <Row className="flex-nowrap" style={{ display: 'flex', flex: '1' }}>
                <Col md="auto" classname="sticky-sideabr" style={{ borderRight: 'solid 1px' }}>
                    <div className="sticky-top">
                        <h1>{params.service}</h1>
                        {status === 'fetching' && <p>Loading resources...</p>}
                        {status === 'error' && <p>Error: {error}</p>}
                        {status === 'fetched' && (
                            <nav sstyle={{ paddingRight: '1rem', display: 'flex', flexDirection: 'column', flex: '1' }}>
                                {data.resources.map(resource => (
                                    <NavLink key={resource} to={`/aws/${params.service}/${resource}`} style={({ isActive }) => {
                                        return {
                                            display: 'block',
                                            margin: '2px 0',
                                            color: isActive ? 'green' : '',
                                        };
                                    }}>
                                        {resource}
                                    </NavLink>
                                ))
                                }
                            </nav>
                        )}
                    </div>
                </Col>
                <Col style={{ overflowY: 'auto' }}>
                    <Outlet />
                </Col>
            </Row>
        </Container>
    )
}