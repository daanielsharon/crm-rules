import React, { useState, useEffect } from 'react';
import { 
  Box, 
  TextField, 
  Typography, 
  CircularProgress,
  Button,
  Tooltip
} from '@mui/material';
import RefreshIcon from '@mui/icons-material/Refresh';
import logService from '../services/logService';
import LogsTable from '../components/LogsTable';

const LogsPage = () => {
  const [logs, setLogs] = useState([]);
  const [loading, setLoading] = useState(false);
  const [searchUserId, setSearchUserId] = useState('');
  const [searchRuleId, setSearchRuleId] = useState('');

  const fetchLogs = async () => {
    setLoading(true);
    try {
      const data = await logService.getLogs(searchUserId, searchRuleId);
      setLogs(data);
    } catch (error) {
      console.error('Error in LogsPage fetchLogs:', error);
      console.error('Error details:', JSON.stringify(error, null, 2));
      setLogs([]);
    } finally {
      setLoading(false);
      console.log('Loading complete, current logs:', logs);
    }
  };

  useEffect(() => {
    fetchLogs();
  }, [searchUserId, searchRuleId]);

  return (
    <Box p={2}>
      <Box mb={3} display="flex" gap={2} alignItems="center">
        <TextField
          label="Filter by User ID (optional)"
          value={searchUserId}
          onChange={(e) => {
            setSearchUserId(e.target.value);
          }}
          variant="outlined"
          size="small"
          fullWidth
          placeholder="Enter user ID to filter"
        />
        <TextField
          label="Filter by Rule ID (optional)"
          value={searchRuleId}
          onChange={(e) => {
            setSearchRuleId(e.target.value);
          }}
          variant="outlined"
          size="small"
          fullWidth
          placeholder="Enter rule ID to filter"
        />
        <Tooltip title="Refresh Logs">
          <Button 
            variant="outlined" 
            color="primary" 
            onClick={fetchLogs} 
            disabled={loading}
            startIcon={<RefreshIcon />}
            size="small"
            sx={{ 
              height: '40px', 
              px: 2,  
              minWidth: '120px'  
            }}
          >
            Refresh
          </Button>
        </Tooltip>
      </Box>

      {loading ? (
        <Box 
          display="flex" 
          justifyContent="center" 
          alignItems="center" 
          height="300px"
        >
          <CircularProgress />
        </Box>
      ) : logs.length === 0 ? (
        <Typography variant="body1" align="center" color="textSecondary">
          {(searchUserId || searchRuleId) 
            ? 'No logs found matching the search criteria' 
            : 'No logs available'}
        </Typography>
      ) : (
        <LogsTable logs={logs} />
      )}
    </Box>
  );
};

export default LogsPage;
