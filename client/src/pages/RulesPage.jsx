import React, { useState, useEffect } from 'react';
import { Box, Button } from '@mui/material';
import AddRuleModal from '../components/AddRuleModal';
import RuleTable from '../components/RulesTable';
import ruleService from '../services/ruleService';

const RulesPage = () => {
  const [rules, setRules] = useState([]);
  const [loading, setLoading] = useState(false);
  const [openModal, setOpenModal] = useState(false);
  const [newRule, setNewRule] = useState({ 
    name: '',
    condition: '', 
    schedule: '',
    actions: [] 
  });

  useEffect(() => {
    fetchRules();
  }, []);

  const fetchRules = async () => {
    setLoading(true);
    try {
      const data = await ruleService.getRules();
      setRules(data);
    } catch (error) {
      console.error('Error fetching rules:', error);
    } finally {
      setLoading(false);
    }
  };

  const handleAddRule = async () => {
    try {
      await ruleService.createRule(newRule);
      await fetchRules();
      setOpenModal(false);
      setNewRule({ 
        name: '',
        condition: '', 
        schedule: '',
        actions: [] 
      });
    } catch (error) {
      console.error('Error adding rule:', error);
    }
  };

  return (
    <Box p={2}>
      <Button 
        variant="contained" 
        color="primary" 
        onClick={() => setOpenModal(true)}
        sx={{ my: 2 }}
      >
        Add Rule
      </Button>
      
      <RuleTable 
        rules={rules} 
        onRuleDelete={() => {
          fetchRules(); 
        }}
        onRuleUpdate={() => {
          fetchRules(); 
        }}
      />

      <AddRuleModal
        open={openModal}
        onClose={() => setOpenModal(false)}
        onSubmit={handleAddRule}
        rule={newRule}
        setRule={setNewRule}
      />
    </Box>
  );
};

export default RulesPage;
