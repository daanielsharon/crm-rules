import axios from '../api/axios';

const actionService = {
  getActionsByRuleId: async (ruleId) => {
    try {
      const response = await axios.get(`/actions?rule_id=${ruleId}`);
      return response.data.data || [];
    } catch (error) {
      console.error(`Error fetching actions for rule ${ruleId}:`, error);
      return [];
    }
  },

  createAction: async (actionData) => {
    try {
      const response = await axios.post('/actions', actionData);
      return response.data;
    } catch (error) {
      console.error('Error creating action:', error);
      throw error;
    }
  },

  deleteAction: async (actionId) => {
    try {
      const response = await axios.delete(`/actions/${actionId}`);
      return response.data;
    } catch (error) {
      console.error(`Error deleting action ${actionId}:`, error);
      throw error;
    }
  }
};

export default actionService;
