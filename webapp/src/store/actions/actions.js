const dispatch = (type, payload = null) => ({ type: type, payload: payload});

export const setInitialState = () => dispatch('INITIAL_STATE');