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
  DialogActions,
  Chip,
  Box
} from '@mui/material';
import CheckCircle from '@mui/icons-material/CheckCircle';
import Cancel from '@mui/icons-material/Cancel';
import userService from '../services/userService';
import AddUserModal from './AddUserModal';

const UsersTable = ({ users, onUserDelete }) => {
  const [openConfirmDialog, setOpenConfirmDialog] = useState(false);
  const [userToDelete, setUserToDelete] = useState(null);
  
  const [userToUpdate, setUserToUpdate] = useState(null);

  const handleDeleteUser = (user) => {
    setUserToDelete(user);
    setOpenConfirmDialog(true);
  };

  const confirmDeleteUser = async () => {
    if (userToDelete) {
      try {
        await userService.deleteUser(userToDelete.id);
        
        if (onUserDelete) {
          onUserDelete();
        }
        
        setOpenConfirmDialog(false);
        setUserToDelete(null);
      } catch (error) {
        console.error('Error deleting user:', error);
      }
    }
  };

  const handleUpdateUser = (user) => {
    setUserToUpdate({
      id: user.id,
      name: user.name,
      email: user.email,
      plan: user.plan,
      failed_logins: user.failed_logins || 0
    });
    setOpenUpdateModal(true);
  };

  const handleUpdateSubmit = async () => {
    try {
      const userToUpdateWithIntFailedLogins = {
        ...userToUpdate,
        failed_logins: parseInt(userToUpdate.failed_logins, 10) || 0
      };

      await userService.updateUser(userToUpdate.id, userToUpdateWithIntFailedLogins);
      
      if (onUserDelete) {
        onUserDelete();
      }
      
      setOpenUpdateModal(false);
      setUserToUpdate(null);
    } catch (error) {
      console.error('Error updating user:', error);
    }
  };

  const handleCloseConfirmDialog = () => {
    setOpenConfirmDialog(false);
    setUserToDelete(null);
  };

  const renderEmailVerified = (verified) => {
    return verified ? (
      <Chip 
        icon={<CheckCircle />} 
        label="Verified" 
        color="success" 
        size="small" 
        variant="outlined"
      />
    ) : (
      <Chip 
        icon={<Cancel />} 
        label="Unverified" 
        color="error" 
        size="small" 
        variant="outlined"
      />
    );
  };

  return (
    <>
      <TableContainer component={Paper}>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>ID</TableCell>
              <TableCell>Name</TableCell>
              <TableCell>Email</TableCell>
              <TableCell>Plan</TableCell>
              <TableCell>Last Active</TableCell>
              <TableCell>Failed Logins</TableCell>
              <TableCell>Email Verified</TableCell>
              <TableCell>Created At</TableCell>
              <TableCell>Updated At</TableCell>
              <TableCell>Actions</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {users.map((user) => (
              <TableRow key={user.id}>
                <TableCell>{user.id}</TableCell>
                <TableCell>{user.name}</TableCell>
                <TableCell>{user.email}</TableCell>
                <TableCell>{user.plan}</TableCell>
                <TableCell>{user.last_active}</TableCell>
                <TableCell>{user.failed_logins || 0}</TableCell>
                <TableCell>
                  {renderEmailVerified(user.email_verified)}
                </TableCell>
                <TableCell>{user.created_at}</TableCell>
                <TableCell>{user.updated_at}</TableCell>
                <TableCell>
                  <Box display="flex" gap={1}>
                    <Button 
                      variant="outlined" 
                      color="primary" 
                      size="small"
                      onClick={() => handleUpdateUser(user)}
                    >
                      Update
                    </Button>
                    <Button 
                      variant="outlined" 
                      color="error" 
                      size="small"
                      onClick={() => handleDeleteUser(user)}
                    >
                      Delete
                    </Button>
                  </Box>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>

      {userToUpdate && (
        <AddUserModal
          open={openUpdateModal}
          onClose={() => setOpenUpdateModal(false)}
          onSubmit={handleUpdateSubmit}
          user={userToUpdate}
          setUser={setUserToUpdate}
        />
      )}

      <Dialog
        open={openConfirmDialog}
        onClose={handleCloseConfirmDialog}
      >
        <DialogTitle>Confirm User Deletion</DialogTitle>
        <DialogContent>
          <DialogContentText>
            Are you sure you want to delete the user {userToDelete?.name}? 
            This action cannot be undone.
          </DialogContentText>
        </DialogContent>
        <DialogActions>
          <Button onClick={handleCloseConfirmDialog} color="primary">
            Cancel
          </Button>
          <Button onClick={confirmDeleteUser} color="error" autoFocus>
            Delete
          </Button>
        </DialogActions>
      </Dialog>
    </>
  );
};

export default UsersTable;
