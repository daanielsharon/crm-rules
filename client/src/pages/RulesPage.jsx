import React, { useState, useEffect } from 'react';
import { Box, Button } from '@mui/material';
import AddRuleModal from '../components/AddRuleModal';
// import { getRules, createRule } from '../services/ruleService';
import RuleTable from '../components/RulesTable';

const RulesPage = () => {
  const [rules, setRules] = useState([]);
  const [loading, setLoading] = useState(false);
  const [openModal, setOpenModal] = useState(false);
  const [newRule, setNewRule] = useState({ condition: '', actions: [] });

  useEffect(() => {
    fetchRules();
  }, []);

  const fetchRules = async () => {
    setLoading(true);
    try {
      const data = await getRules();
      setRules(data);
    } catch (error) {
      console.error('Error fetching rules:', error);
    } finally {
      setLoading(false);
    }
  };

  const handleAddRule = async () => {
    try {
      await createRule(newRule);
      fetchRules();
      setOpenModal(false);
      setNewRule({ condition: '', actions: [] });
    } catch (error) {
      console.error('Error adding rule:', error);
    }
  };

  return (
    <Box p={2}>
      <Button variant="contained" color="primary" onClick={() => setOpenModal(true)}>
        Add Rule
      </Button>
      <RuleTable rules={rules} />
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
