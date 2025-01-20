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
import ruleService from '../services/ruleService';
import actionService from '../services/actionService';
import ActionsModal from './ActionsModal';

const RulesTable = ({ rules, onRuleDelete }) => {
  const [selectedRuleId, setSelectedRuleId] = useState(null);
  const [openActionsModal, setOpenActionsModal] = useState(false);
  const [openConfirmDialog, setOpenConfirmDialog] = useState(false);
  const [ruleToDelete, setRuleToDelete] = useState(null);
  const [selectedRuleActions, setSelectedRuleActions] = useState([]);
  const [loadingActions, setLoadingActions] = useState(false);

  const handleViewActions = async (ruleId) => {
    setSelectedRuleId(ruleId);
    setLoadingActions(true);
    try {
      const actions = await actionService.getActionsByRuleId(ruleId);
      setSelectedRuleActions(actions);
      setOpenActionsModal(true);
    } catch (error) {
      console.error('Error fetching actions:', error);
    } finally {
      setLoadingActions(false);
    }
  };

  const handleCloseActionsModal = () => {
    setOpenActionsModal(false);
    setSelectedRuleId(null);
    setSelectedRuleActions([]);
  };

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

      <ActionsModal 
        open={openActionsModal}
        onClose={handleCloseActionsModal}
        ruleId={selectedRuleId}
        actions={selectedRuleActions}
        loading={loadingActions}
      />

      <Dialog
        open={openConfirmDialog}
        onClose={() => setOpenConfirmDialog(false)}
      >
        <DialogTitle>Confirm Rule Deletion</DialogTitle>
        <DialogContent>
          <DialogContentText>
            Are you sure you want to delete the rule {ruleToDelete?.name}? 
            This action cannot be undone and will also remove all associated actions.
          </DialogContentText>
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setOpenConfirmDialog(false)} color="primary">
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
