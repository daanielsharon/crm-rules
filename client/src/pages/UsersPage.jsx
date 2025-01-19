import React, { useState, useEffect } from 'react';
import { Button, Box, CircularProgress, Typography } from '@mui/material';
import AddUserModal from '../components/AddUserModal.jsx';
import UserTable from '../components/UsersTable.jsx';
import UserService from '../services/userService.js';

const UsersPage = () => {
  const [users, setUsers] = useState([]);
  const [loading, setLoading] = useState(true);
  const [openModal, setOpenModal] = useState(false);
  const [newUser, setNewUser] = useState({ name: '', email: '' });

  useEffect(() => {
    fetchUsers();
  }, []);

  const fetchUsers = async () => {
    setLoading(true);
    try {
      const data = await UserService.getUsers();
      console.log('data', data)
      setUsers(data);
    } catch (error) {
      console.error('Error fetching users:', error);
    } finally {
      setLoading(false);
    }
  };

  const handleAddUser = async () => {
    try {
      await UserService.createUser(newUser);
      fetchUsers();
      setOpenModal(false);
      setNewUser({ name: '', email: '' });
    } catch (error) {
      console.error('Error adding user:', error);
    }
  };

  return (
    <Box p={2}>
      <Button 
        variant="contained" 
        color="primary" 
        onClick={() => setOpenModal(true)}
      >
        Add User
      </Button>

      {loading ? (
        <Box 
          display="flex" 
          justifyContent="center" 
          alignItems="center" 
          height="200px"
        >
          <CircularProgress />
        </Box>
      ) : users.length === 0 ? (
        <Typography variant="body1" align="center" sx={{ mt: 4 }}>
          No users found. Add a new user to get started.
        </Typography>
      ) : (
        <UserTable 
          users={users} 
          onUserDelete={fetchUsers}
        />
      )}

      <AddUserModal
        open={openModal}
        onClose={() => setOpenModal(false)}
        onSubmit={handleAddUser}
        user={newUser}
        setUser={setNewUser}
      />
    </Box>
  );
};

export default UsersPage;
