import React, { useState } from 'react';
import { 
  Modal, 
  Box, 
  TextField, 
  Button, 
  Typography 
} from '@mui/material';

const AddUserModal = ({ open, onClose, onSubmit, user, setUser }) => {
  const [errors, setErrors] = useState({
    name: '',
    email: '',
    failed_logins: ''
  });

  const validateField = (name, value) => {
    switch(name) {
      case 'name':
        return value.trim() === '' ? 'Name is required' : '';
      case 'email':
        const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        return !emailRegex.test(value) ? 'Invalid email format' : '';
      case 'failed_logins':
        const parsedValue = parseInt(value, 10);
        return (isNaN(parsedValue) || parsedValue < 0) 
          ? 'Failed logins must be a non-negative number' 
          : '';
      default:
        return '';
    }
  };

  const handleChange = (e) => {
    const { name, value } = e.target;
    
    // Special handling for failed_logins to ensure numeric input
    const processedValue = name === 'failed_logins' 
      ? value.replace(/[^0-9]/g, '') 
      : value;
    
    setUser(prev => ({ ...prev, [name]: processedValue }));
    
    const errorMessage = validateField(name, processedValue);
    setErrors(prev => ({ ...prev, [name]: errorMessage }));
  };

  const handleSubmit = () => {
    const nameError = validateField('name', user.name);
    const emailError = validateField('email', user.email);
    const failedLoginsError = validateField('failed_logins', user.failed_logins);

    setErrors({
      name: nameError,
      email: emailError,
      failed_logins: failedLoginsError
    });

    // Only submit if no errors
    if (!nameError && !emailError && !failedLoginsError) {
      onSubmit();
    }
  };

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
        <Typography variant="h6" gutterBottom>
          Add/Update User
        </Typography>
        <TextField
          name="name"
          label="Name"
          fullWidth
          margin="normal"
          value={user.name}
          onChange={handleChange}
          error={!!errors.name}
          helperText={errors.name}
        />
        <TextField
          name="email"
          label="Email"
          fullWidth
          margin="normal"
          value={user.email}
          onChange={handleChange}
          error={!!errors.email}
          helperText={errors.email}
        />
        <TextField
          name="plan"
          label="Plan"
          fullWidth
          margin="normal"
          value={user.plan}
          placeholder="Select plan (free,premium)"
          onChange={handleChange}
        />
        <TextField
          name="failed_logins"
          label="Failed Logins"
          fullWidth
          margin="normal"
          type="number"
          value={user.failed_logins}
          onChange={handleChange}
          inputProps={{ min: 0 }}
          error={!!errors.failed_logins}
          helperText={errors.failed_logins}
        />
        <Box mt={2} display="flex" justifyContent="flex-end" gap={2}>
          <Button onClick={onClose} variant="outlined">
            Cancel
          </Button>
          <Button 
            onClick={handleSubmit} 
            variant="contained" 
            color="primary"
          >
            Submit
          </Button>
        </Box>
      </Box>
    </Modal>
  );
};

export default AddUserModal;
