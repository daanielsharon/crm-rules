import React, { useState, useEffect } from 'react';
import { Box, TableContainer, Table, TableHead, TableRow, TableCell, TableBody, TextField } from '@mui/material';
// import { getLogs } from '../services/logService';

const LogsPage = () => {
  const [logs, setLogs] = useState([]);
  const [filteredLogs, setFilteredLogs] = useState([]);
  const [loading, setLoading] = useState(false);
  const [searchRuleId, setSearchRuleId] = useState('');
  const [searchUserId, setSearchUserId] = useState('');

  useEffect(() => {
    fetchLogs();
  }, []);

  const fetchLogs = async () => {
    setLoading(true);
    try {
      const data = await getLogs();
      setLogs(data);
      setFilteredLogs(data); // Initially display all logs
    } catch (error) {
      console.error('Error fetching logs:', error);
    } finally {
      setLoading(false);
    }
  };

  const handleFilter = () => {
    let filtered = logs;
    if (searchRuleId) {
      filtered = filtered.filter((log) => log.rule_id.includes(searchRuleId));
    }
    if (searchUserId) {
      filtered = filtered.filter((log) => log.user_id.includes(searchUserId));
    }
    setFilteredLogs(filtered);
  };

  return (
    <Box p={2}>
      <Box mb={3} display="flex" gap={2}>
        <TextField
          label="Filter by Rule ID"
          value={searchRuleId}
          onChange={(e) => setSearchRuleId(e.target.value)}
          onBlur={handleFilter}
        />
        <TextField
          label="Filter by User ID"
          value={searchUserId}
          onChange={(e) => setSearchUserId(e.target.value)}
          onBlur={handleFilter}
        />
      </Box>

      <TableContainer>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>ID</TableCell>
              <TableCell>Rule ID</TableCell>
              <TableCell>User ID</TableCell>
              <TableCell>Action</TableCell>
              <TableCell>Status</TableCell>
              <TableCell>Executed At</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {filteredLogs.map((log) => (
              <TableRow key={log.id}>
                <TableCell>{log.id}</TableCell>
                <TableCell>{log.rule_id}</TableCell>
                <TableCell>{log.user_id}</TableCell>
                <TableCell>{log.action}</TableCell>
                <TableCell>{log.status}</TableCell>
                <TableCell>{log.executed_at}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
    </Box>
  );
};

export default LogsPage;
