import React, { useState } from 'react';

import 'semantic-ui-css/semantic.min.css'
import { Button, ButtonGroup } from 'semantic-ui-react';

const Home = () => {
    return (
        <>
            <h1>Welcome!</h1>
            <h2>What would you like to do?</h2>
            <div class="ui button" tabindex="0">View Experiments</div>
            <div class="ui button" tabindex="0">View Participants</div>
            <div class="ui button" tabindex="0">Rejected Users</div>
            <div class="ui button" tabindex="0">View Sign Ups</div>
        </>
    )
};

export default Home;