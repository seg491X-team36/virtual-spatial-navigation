import React, { useState } from 'react';

import 'semantic-ui-css/semantic.min.css'
import { Button, ButtonGroup } from 'semantic-ui-react';

const Home = () => {
    return (
        <>
            <h1>Welcome!</h1>
            <h2>What would you like to do?</h2>
            <button class="ui button" tabindex="0">View Experiments</button>
            <button class="ui button" tabindex="0">View Participants</button>
            <button class="ui button" tabindex="0">Rejected Users</button>
            <button class="ui button" tabindex="0">View Sign Ups</button>
        </>
    )
};

export default Home;