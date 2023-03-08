import 'semantic-ui-css/semantic.min.css'
import { Button, Checkbox, Icon, Table } from 'semantic-ui-react'

import NavBar from '../components/NavBar';

const ViewSignUps = () => {
    return (
        <>
            <h1>Sign Ups</h1>
            <p>Here are participants who have signed up for the trial. Please select the ones you would like to accept.</p>
            <Table compact celled>
                <Table.Header>
                    <Table.Row>
                        <Table.HeaderCell>Accept</Table.HeaderCell>
                        <Table.HeaderCell>Name</Table.HeaderCell>
                        <Table.HeaderCell>Email</Table.HeaderCell>
                    </Table.Row>
                </Table.Header>

                <Table.Body>
                    <Table.Row>
                        <Table.Cell collapsing>
                            <Checkbox ui checkbox />
                        </Table.Cell>
                        <Table.Cell>John Smith</Table.Cell>
                        <Table.Cell>jsmit345@uottawa.ca</Table.Cell>
                    </Table.Row>
                    <Table.Row>
                        <Table.Cell collapsing>
                            <Checkbox ui checkbox />
                        </Table.Cell>
                        <Table.Cell>Oisin Gallagher</Table.Cell>
                        <Table.Cell>ogall365@uottawa.ca</Table.Cell>
                    </Table.Row>
                    <Table.Row>
                        <Table.Cell collapsing>
                            <Checkbox ui checkbox />
                        </Table.Cell>
                        <Table.Cell>Carmen Wagner</Table.Cell>
                        <Table.Cell>cwagn124@uottawa.ca</Table.Cell>
                    </Table.Row>
                </Table.Body>

                <Table.Footer fullWidth>
                    <Table.Row>
                        <Table.HeaderCell />
                        <Table.HeaderCell colSpan='4'>
                            <Button
                                floated='right'
                                icon
                                labelPosition='left'
                                primary
                                size='small'
                            >
                                <Icon name='user plus' /> Add Participant
                            </Button>
                            <Button size='small' class="ui primary button">Accept</Button>
                            <Button size='small'>
                                Accept All
                            </Button>
                        </Table.HeaderCell>
                    </Table.Row>
                </Table.Footer>
            </Table>
        </>
    )
};

export default ViewSignUps;