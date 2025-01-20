import axios from '../api/axios';

const logService = {
  getLogs: async (userId = '', ruleId = '') => {
    const params = new URLSearchParams();
    if (userId) params.append('user_id', userId);
    if (ruleId) params.append('rule_id', ruleId);
    const response = await axios.get(`/logs${params.toString() ? `?${params.toString()}` : ''}`);
    return response.data.data || [];
  },
};

export default logService;
