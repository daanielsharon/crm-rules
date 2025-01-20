import React from 'react';
import { 
  TableContainer, 
  Table, 
  TableHead, 
  TableRow, 
  TableCell, 
  TableBody, 
  Paper,
  Chip
} from '@mui/material';

const LogsTable = ({ logs }) => {
  const renderStatusChip = (status) => {
    const colorMap = {
      'success': 'success',
      'failed': 'error',
      'pending': 'warning',
      'skipped': 'default'
    };

    return (
      <Chip 
        label={status} 
        color={colorMap[status.toLowerCase()] || 'default'} 
        size="small" 
        variant="outlined"
      />
    );
  };

  return (
    <TableContainer component={Paper}>
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
          {logs.map((log) => (
            <TableRow key={log.id}>
              <TableCell>{log.id}</TableCell>
              <TableCell>{log.rule_id}</TableCell>
              <TableCell>{log.user_id}</TableCell>
              <TableCell>{log.action}</TableCell>
              <TableCell>
                {renderStatusChip(log.status)}
              </TableCell>
              <TableCell>{log.executed_at}</TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
};

export default LogsTable;
