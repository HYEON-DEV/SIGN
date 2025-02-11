// API 호출 모듈

import axios from 'axios';

// const API_BASE_URL = 

export const createDID = async () => {
    return await axios.post('/api/did/create')
};

export const fetchVC = async () => {
    return await axios.get('/api/vc/list')
};

export const createVP = async () => {
    return await axios.post('/api/vp/create', {vc});
};