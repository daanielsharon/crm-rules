import React from 'react';
import { Modal, Box, TextField, Button, Select, MenuItem, InputLabel, FormControl } from '@mui/material';

const AddRuleModal = ({ open, onClose, onSubmit, rule, setRule }) => {
  return (
    <Modal open={open} onClose={onClose}>
      <Box
        sx={{
          position: 'absolute',
          top: '50%',
          left: '50%',
          transform: 'translate(-50%, -50%)',
          width: 400,
          bgcolor: 'background.paper',
          p: 4,
          borderRadius: '8px',
        }}
      >
        <h2>Add Rule</h2>
        <TextField
          label="Condition"
          fullWidth
          margin="normal"
          value={rule.condition}
          onChange={(e) => setRule({ ...rule, condition: e.target.value })}
        />
        <FormControl fullWidth margin="normal">
          <InputLabel>Actions</InputLabel>
          <Select
            multiple
            value={rule.actions}
            onChange={(e) => setRule({ ...rule, actions: e.target.value })}
            renderValue={(selected) => selected.join(', ')}
          >
            <MenuItem value="send_notification">Send Notification</MenuItem>
            <MenuItem value="send_email">Send Email</MenuItem>
          </Select>
        </FormControl>
        <Box mt={2} display="flex" justifyContent="flex-end" gap={2}>
          <Button onClick={onClose} variant="outlined">
            Cancel
          </Button>
          <Button onClick={onSubmit} variant="contained" color="primary">
            Submit
          </Button>
        </Box>
      </Box>
    </Modal>
  );
};

export default AddRuleModal;
