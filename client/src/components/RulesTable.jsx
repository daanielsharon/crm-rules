import React, { useState } from 'react';
import { 
  Table, 
  TableBody, 
  TableCell, 
  TableContainer, 
  TableHead, 
  TableRow, 
  Paper,
  Button,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogContentText,
  DialogActions
} from '@mui/material';
import actionService from '../services/actionService';
import ruleService from '../services/ruleService';
import ActionModal from './ActionModal';

const RulesTable = ({ rules, onRuleDelete }) => {
  const [selectedRuleActions, setSelectedRuleActions] = useState([]);
  const [isActionsModalOpen, setIsActionsModalOpen] = useState(false);
  const [selectedRuleId, setSelectedRuleId] = useState(null);
  
  const [ruleToDelete, setRuleToDelete] = useState(null);

  const handleViewActions = async (ruleId) => {
    try {
      const actions = await actionService.getActionsByRuleId(ruleId);
      setSelectedRuleActions(actions);
      setSelectedRuleId(ruleId);
      setIsActionsModalOpen(true);
    } catch (error) {
      console.error('Error fetching actions:', error);
    }
  };

  const handleCloseActionsModal = () => {
    setIsActionsModalOpen(false);
    setSelectedRuleActions([]);
    setSelectedRuleId(null);
  }

  const handleDeleteRule = (rule) => {
    setRuleToDelete(rule);
    setOpenConfirmDialog(true);
  };

  const confirmDeleteRule = async () => {
    if (ruleToDelete) {
      try {
        await ruleService.deleteRule(ruleToDelete.id);
        
        if (onRuleDelete) {
          onRuleDelete(ruleToDelete.id);
        }
        
        setOpenConfirmDialog(false);
        setRuleToDelete(null);
      } catch (error) {
        console.error('Error deleting rule:', error);
      }
    }
  };

  const handleCloseConfirmDialog = () => {
    setOpenConfirmDialog(false);
    setRuleToDelete(null);
  };

  return (
    <>
      <TableContainer component={Paper}>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>ID</TableCell>
              <TableCell>Name</TableCell>
              <TableCell>Condition</TableCell>
              <TableCell>Schedule</TableCell>
              <TableCell>Actions</TableCell>
              <TableCell>Operations</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {rules.map((rule) => (
              <TableRow key={rule.id}>
                <TableCell>{rule.id}</TableCell>
                <TableCell>{rule.name}</TableCell>
                <TableCell>{rule.condition}</TableCell>
                <TableCell>{rule.schedule}</TableCell>
                <TableCell>
                  <Button 
                    variant="outlined" 
                    color="primary" 
                    size="small"
                    onClick={() => handleViewActions(rule.id)}
                  >
                    View Actions
                  </Button>
                </TableCell>
                <TableCell>
                  <Button 
                    variant="outlined" 
                    color="error" 
                    size="small"
                    onClick={() => handleDeleteRule(rule)}
                  >
                    Delete
                  </Button>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>

      <ActionModal 
        open={isActionsModalOpen}
        onClose={handleCloseActionsModal}
        actions={selectedRuleActions}
        ruleId={selectedRuleId}
      />

      <Dialog
        open={openConfirmDialog}
        onClose={handleCloseConfirmDialog}
      >
        <DialogTitle>Confirm Rule Deletion</DialogTitle>
        <DialogContent>
          <DialogContentText>
            Are you sure you want to delete the rule {ruleToDelete?.name}? 
            This action cannot be undone and will also remove all associated actions.
          </DialogContentText>
        </DialogContent>
        <DialogActions>
          <Button onClick={handleCloseConfirmDialog} color="primary">
            Cancel
          </Button>
          <Button onClick={confirmDeleteRule} color="error" autoFocus>
            Delete
          </Button>
        </DialogActions>
      </Dialog>
    </>
  );
};

export default RulesTable;
