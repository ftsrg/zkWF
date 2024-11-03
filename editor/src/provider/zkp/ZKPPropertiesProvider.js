// Import your custom property entries.
// The entry is a text input field with logic attached to create,
// update and delete the "spell" property.
import PublicKeyProps from './parts/PublicKeyProps';
import VariableProps from './parts/VariableProps';
import ParticipantProps from './parts/PaymentTaskProps';

import { is } from 'bpmn-js/lib/util/ModelUtil';

const LOW_PRIORITY = 500;


/**
 * A provider with a `#getGroups(element)` method
 * that exposes groups for a diagram element.
 *
 * @param {PropertiesPanel} propertiesPanel
 * @param {Function} translate
 */
export default function ZKPPropertiesProvider(propertiesPanel, translate) {

  // API ////////

  /**
   * Return the groups provided for the given element.
   *
   * @param {DiagramElement} element
   *
   * @return {(Object[]) => (Object[])} groups middleware
   */
  this.getGroups = function(element) {

    /**
     * We return a middleware that modifies
     * the existing groups.
     *
     * @param {Object[]} groups
     *
     * @return {Object[]} modified groups
     */
    return function(groups) {

      // Add the "magic" group
      if(is(element, 'bpmn:Task')) {
        groups.push(createZKPGroup(element, translate));
      }
      if(is(element, 'bpmn:Lane') || is(element, 'bpmn:Participant')) {
        groups.push(createZKPGroup2(element, translate));
      }
      if(is(element, 'bpmn:Task') && element.businessObject.$attrs.type === 'paymentTask') {
        groups.push(createPaymentGroup(element, translate));
      }

      return groups;
    }
  };


  // registration ////////

  // Register our custom magic properties provider.
  // Use a lower priority to ensure it is loaded after
  // the basic BPMN properties.
  propertiesPanel.registerProvider(LOW_PRIORITY, this);
}

ZKPPropertiesProvider.$inject = [ 'propertiesPanel', 'translate' ];

// Create the custom magic group
function createZKPGroup(element, translate) {

  // create a group called "Magic properties".
  const zkpGroup = {
    id: 'zkp',
    label: translate('ZKP properties'),
    entries: [
     ...VariableProps(element)]
  };

  return zkpGroup
}


// Create the custom magic group
function createZKPGroup2(element, translate) {

  // create a group called "Magic properties".
  const zkpGroup = {
    id: 'zkp',
    label: translate('ZKP properties'),
    entries: [
     ...PublicKeyProps( element )]
  };

  return zkpGroup
}

function createPaymentGroup(element, translate) {
  const paymentGroup = {
    id: 'payment',
    label: translate('Payment properties'),
    entries: [
      ...ParticipantProps(element)
    ]
  };

  return paymentGroup;
}


