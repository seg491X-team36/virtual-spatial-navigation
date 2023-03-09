import React, {useState} from 'react';
import 'semantic-ui-css/semantic.min.css'

import * as fake_data from '../components/apiFunctions.js';
import {Header, Grid, Container, Menu} from 'semantic-ui-react';
import ExperimentDetails from '../components/ExperimentDetails.js'

const ViewExperiments = () => {

    const experiments = fake_data.experiments();
    const[selectedExperiment, setSelectedExperiment] = useState({}); // selectedExperiment is an experiment obj

    const handleViewExperimentClick = (e) => {
        setSelectedExperiment({
            id: e.id,
            name: e.name,
            description: e.description,
            config: e.config,
            results: e.results,
            pending: e.pending,
            usersNotInvited: e.usersNotInvited
        });
    }
    
    return(
        <Container>
        <h1>Experiments</h1>
        <Grid columns='equal'>
            <Grid.Column>
                <Menu vertical>
                {experiments.length > 0 ? experiments.map(exp => {return (
                <Menu.Item
                    name={exp.name}
                    onClick={() => handleViewExperimentClick(exp)}
                    active={selectedExperiment.name === exp.name}>
                    <Header as='h4'>{exp.name}</Header>
                    <p>{exp.description}</p>
                </Menu.Item>
                )})
                :
                <p>There are no experiments to display.</p>
                }
                </Menu>
            </Grid.Column>

            <Grid.Column width={12}>
            {selectedExperiment.id !== undefined ?
                <ExperimentDetails exp_id={selectedExperiment.id}/>
                :
                <></>
                }
            </Grid.Column>
        </Grid>
        </Container>
    )
};

export default ViewExperiments;