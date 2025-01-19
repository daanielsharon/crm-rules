import axios from '../api/axios';

const logService = {
  getLogs: async (filters = {}) => {
    try {
      const response = await axios.get('/logs', { params: filters });
      return response.data;
    } catch (error) {
      console.error('Error fetching logs:', error);
      throw error;
    }
  },

  getLogById: async (logId) => {
    try {
      const response = await axios.get(`/logs/${logId}`);
      return response.data;
    } catch (error) {
      console.error(`Error fetching log ${logId}:`, error);
      throw error;
    }
  },

  deleteLogs: async (logIds) => {
    try {
      const response = await axios.delete('/logs', { 
        data: { logIds } 
      });
      return response.data;
    } catch (error) {
      console.error('Error deleting logs:', error);
      throw error;
    }
  }
};

export default logService;
