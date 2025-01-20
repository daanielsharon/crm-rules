import React, { useState, useEffect } from 'react';
import { 
  Dialog, 
  DialogTitle, 
  DialogContent, 
  TextField, 
  Button, 
  Table, 
  TableBody, 
  TableCell, 
  TableContainer, 
  TableHead, 
  TableRow, 
  Paper,
  Box,
  CircularProgress,
  Typography,
  Alert
} from '@mui/material';
import actionService from '../services/actionService';

const ActionsModal = ({ 
  open, 
  onClose, 
  ruleId, 
  actions: initialActions = [], 
  loading: initialLoading = false 
}) => {
  const [actions, setActions] = useState([]);
  const [newActionName, setNewActionName] = useState('');
  const [error, setError] = useState(null);

  useEffect(() => {
    // Reset state when modal opens
    if (open) {
      setError(null);
      setActions(initialActions);
      setNewActionName('');
    }
  }, [open, initialActions]);

  const handleAddAction = async () => {
    if (!newActionName.trim()) return;

    try {
      await actionService.createAction({
        action: newActionName,
        rule_id: ruleId
      });
      
      const updatedActions = await actionService.getActionsByRuleId(ruleId);
      setActions(updatedActions);
      
      setNewActionName('');
      setError(null);
    } catch (error) {
      console.error('Error adding action:', error);
      setError(error.message || 'Failed to add action');
    }
  };

  const formatDate = (dateString) => {
    if (!dateString) return 'N/A';
    try {
      return new Date(dateString).toLocaleString();
    } catch {
      return 'Invalid Date';
    }
  };

  return (
    <Dialog open={open} onClose={onClose} maxWidth="md" fullWidth>
      <DialogTitle>Actions for Rule</DialogTitle>
      <DialogContent>
        {error && (
          <Box mb={2}>
            <Alert severity="error">{error}</Alert>
          </Box>
        )}

        <Box display="flex" gap={2} mb={2}>
          <TextField
            label="New Action Name"
            value={newActionName}
            onChange={(e) => setNewActionName(e.target.value)}
            fullWidth
            disabled={initialLoading}
          />
          <Button 
            variant="contained" 
            color="primary" 
            onClick={handleAddAction}
            disabled={!newActionName.trim() || initialLoading}
          >
            Add Action
          </Button>
        </Box>

        {initialLoading ? (
          <Box display="flex" justifyContent="center" p={3}>
            <CircularProgress />
          </Box>
        ) : actions.length === 0 ? (
          <Box textAlign="center" p={3}>
            <Typography color="textSecondary">
              No actions found for this rule
            </Typography>
          </Box>
        ) : (
          <TableContainer component={Paper}>
            <Table>
              <TableHead>
                <TableRow>
                  <TableCell>ID</TableCell>
                  <TableCell>Action</TableCell>
                  <TableCell>Created At</TableCell>
                  <TableCell>Updated At</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {actions.map((action) => (
                  <TableRow key={action.id}>
                    <TableCell>{action.id}</TableCell>
                    <TableCell>{action.action}</TableCell>
                    <TableCell>{formatDate(action.created_at)}</TableCell>
                    <TableCell>{formatDate(action.updated_at)}</TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </TableContainer>
        )}
      </DialogContent>
      <Box display="flex" justifyContent="flex-end" p={2}>
        <Button onClick={onClose} color="primary" disabled={initialLoading}>
          Close
        </Button>
      </Box>
    </Dialog>
  );
};

export default ActionsModal;
