// these functions can be kept for hooking onto the backend ////////////////////////////////
export const login = (email, password) => {
    return true;
}

export const user = (id) => {
    return data_users.filter(u => u.id === id)[0];
}

export const users = (state) => {
    return data_users.filter(u => u.state === state);
}

export const invite = (id) => {
    return data_invites.filter(i => i.id === id)[0];
}

export const experiment = (id) => {
    return data_experiments.filter(e => e.id === id)[0];
}

export const experiments = () => {
    var res = [];
    data_experiments.forEach(exp => {
        res.push({id: exp.id, name: exp.name, description: exp.description})
    });
    return res;
}

// helper functions, can be deleted later
export const getUserInvites = (u_id) => {
    return data_invites.filter(i => i.user === u_id);
}

export const getUserResults = (u_id) => {
    return data_experiment_results.filter(e => e.user === u_id);
}

export const getExperimentResults = (exp_id) => {
    return data_experiment_results.filter(res => res.id === exp_id);
}

export const getPendingInvites = (exp_id) => {
    return data_invites.filter(invite => data_experiment_results.indexOf(invite.experiment_id) < 0 && invite.experiment === exp_id);
}

export const getUsersNotInvited = (exp_id) => {
    return data_users.filter(u => (u.invites.filter(i => i.experiment === exp_id)).length === 0 && u.state === 'ACCEPTED');
}

// fake data, remove later ////////////////////////////////
export const data_experiment_results = [
    {id: "1", created_at: "", user: "4", experiment: "1"},
    {id: "2", created_at: "", user: "5", experiment: "1"},
    {id: "3", created_at: "", user: "5", experiment: "2"}
];

export const data_invites = [
    {id:"1", created_at: "", user: "4", experiment: "1"},
    {id:"3", created_at: "", user: "5", experiment: "1"}
];

export const data_users = [
    {id: '1', email: 'user1@uottawa.ca', state: 'REGISTERED', source: 'GOOGLEFORMS', invites: getUserInvites('1'), results: getUserResults('1')},
    {id: '2', email: 'user2@uottawa.ca', state: 'REGISTERED', source: 'GOOGLEFORMS', invites: getUserInvites('2'), results: getUserResults('2')},
    {id: '3', email: 'user3@uottawa.ca', state: 'REGISTERED', source: 'GOOGLEFORMS', invites: getUserInvites('3'), results: getUserResults('3')},
    {id: '4', email: 'user4@uottawa.ca', state: 'ACCEPTED', source: 'GOOGLEFORMS', invites: getUserInvites('4'), results: getUserResults('4')},
    {id: '5', email: 'user5@uottawa.ca', state: 'ACCEPTED', source: 'MANUAL', invites: getUserInvites('5'), results: getUserResults('5')},
    {id: '6', email: 'user6@uottawa.ca', state: 'REJECTED', source: 'GOOGLEFORMS', invites: getUserInvites('6'), results: getUserResults('6')},
    {id: '7', email: 'user7@uottawa.ca', state: 'ACCEPTED', source: 'MANUAL', invites: getUserInvites('7'), results: getUserResults('7')}
];

export const data_experiments = [
    {id: "1", name: "Experiment 1", description: 'description', config: {rounds: 5, resume: 'RESET_ROUND'}, results: getExperimentResults(), pending: getPendingInvites('1'), usersNotInvited: getUsersNotInvited('1')},
    {id: "2", name: "Experiment 2", description: 'description', config: {rounds: 3, resume: 'CONTINUE_ROUND'}, results: getExperimentResults(), pending: getPendingInvites('2'), usersNotInvited: getUsersNotInvited('2')}
];