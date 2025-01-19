import React from 'react';
import { 
  Dialog, 
  DialogTitle, 
  DialogContent, 
  List, 
  ListItem, 
  ListItemText,
  Typography,
  Box
} from '@mui/material';

const ActionModal = ({ 
  open, 
  onClose, 
  actions = [], 
  ruleId 
}) => {
  return (
    <Dialog 
      open={open} 
      onClose={onClose}
      maxWidth="md"
      fullWidth
    >
      <DialogTitle>
        <Box display="flex" justifyContent="space-between" alignItems="center">
          <Typography variant="h6">
            Actions for Rule {ruleId}
          </Typography>
        </Box>
      </DialogTitle>
      <DialogContent>
        {actions.length > 0 ? (
          <List>
            {actions.map((action, index) => (
              <ListItem key={index}>
                <ListItemText 
                  primary={action.action} 
                  secondary={`Created: ${new Date(action.created_at).toLocaleString()}`}
                />
              </ListItem>
            ))}
          </List>
        ) : (
          <Typography variant="body1" color="textSecondary" align="center">
            No actions found for this rule
          </Typography>
        )}
      </DialogContent>
    </Dialog>
  );
};

export default ActionModal;
