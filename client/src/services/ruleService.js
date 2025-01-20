import axios from '../api/axios';

const ruleService = {
  getRules: async () => {
    try {
      const response = await axios.get('/rules');
      return response.data.data || [];
    } catch (error) {
      console.error('Error fetching rules:', error);
      throw error;
    }
  },

  getRuleById: async (ruleId) => {
    try {
      const response = await axios.get(`/rules/${ruleId}`);
      return response.data.data;
    } catch (error) {
      console.error(`Error fetching rule ${ruleId}:`, error);
      throw error;
    }
  },

  createRule: async (ruleData) => {
    try {
      const response = await axios.post('/rules', ruleData);
      return response.data.data;
    } catch (error) {
      console.error('Error creating rule:', error);
      throw error;
    }
  },

  updateRule: async (ruleId, ruleData) => {
    try {
      const response = await axios.put(`/rules/${ruleId}`, ruleData);
      return response.data.data;
    } catch (error) {
      console.error(`Error updating rule ${ruleId}:`, error);
      throw error;
    }
  },

  deleteRule: async (ruleId) => {
    try {
      const response = await axios.delete(`/rules/${ruleId}`);
      return response.data;
    } catch (error) {
      console.error(`Error deleting rule ${ruleId}:`, error);
      throw error;
    }
  } 
};

export default ruleService;