import React from 'react';
import { Outlet } from 'react-router-dom';
import { Navbar, Nav, Container } from 'react-bootstrap';

export default function HomePage() {
    return (
        <Container fluid className="bg-opacity-50 vh-100" style={{ "maxHeight": "100vh", display: 'flex', flexDirection: 'column' }}>
            <Navbar bg="light" expand="lg">
                <Navbar.Brand href="/">Cloud Inventory</Navbar.Brand>
                <Navbar.Toggle aria-controls="basic-navbar-nav" />
                <Navbar.Collapse id="basic-navbar-nav">
                    <Nav className="mr-auto">
                        <Nav.Link href="/aws">AWS</Nav.Link>
                    </Nav>
                </Navbar.Collapse>
            </Navbar>
            <Outlet />
        </Container>
    )
}