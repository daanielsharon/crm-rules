import React, { useState } from 'react';
import { Tabs, Tab, Box } from '@mui/material';
import UsersPage from './pages/UsersPage.jsx';
import RulesPage from './pages/RulesPage.jsx';
import LogsPage from './pages/LogsPage.jsx';

const App = () => {
  const [activeTab, setActiveTab] = useState('users');

  const handleChange = (event, newValue) => {
    setActiveTab(newValue);
  };

  return (
    <Box p={2}>
      <h1>Dashboard</h1>
      <Tabs
        value={activeTab}
        onChange={handleChange}
        textColor="primary"
        indicatorColor="primary"
        variant="fullWidth"
        centered
      >
        <Tab value="users" label="Users" />
        <Tab value="rules" label="Rules" />
        <Tab value="logs" label="Logs" />
      </Tabs>

      <Box mt={3}>
        {activeTab === 'users' && <UsersPage />}
        {activeTab === 'rules' && <RulesPage />}
        {activeTab === 'logs' && <LogsPage />}
      </Box>
    </Box>
  );
};

export default App;
