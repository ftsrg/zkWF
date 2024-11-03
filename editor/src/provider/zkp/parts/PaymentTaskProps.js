import { html } from 'htm/preact';

import { SelectEntry, isSelectEntryEdited, NumberFieldEntry, isNumberFieldEntryEdited } from '@bpmn-io/properties-panel';
import { useService } from 'bpmn-js-properties-panel';

// import hooks from the vendored preact package
import { useEffect, useState } from '@bpmn-io/properties-panel/preact/hooks';


import { getBusinessObject, is } from 'bpmn-js/lib/util/ModelUtil';

export default function(element) {

  return [
    {
      id: 'spell',
      element,
      component: ParticipantDropDown,
      isEdited: isSelectEntryEdited
    },
    {
      id: 'amount',
      element,
      component: AmountField,
      isEdited: isNumberFieldEntryEdited
    }
  ];
}

function AmountField(props) {
  const { element, id } = props;

  const modeling = useService('modeling');
  const translate = useService('translate');
  const debounce = useService('debounceInput');
;

  const getValue = () => {
    return element.businessObject.amount || '';
  };

  const setValue = value => {
    return modeling.updateProperties(element, {
      amount: value
    });
  };

  return html`<${NumberFieldEntry}
    id=${ id }
    element=${ element }
    description=${ translate('The amount of money to be transferred') }
    label=${ translate('Amount') }
    getValue=${ getValue }
    setValue=${ setValue }
    debounce=${ debounce }
  />`;
}


function ParticipantDropDown(props) {
  const { element, id } = props;

  const modeling = useService('modeling');
  const translate = useService('translate');
  const debounce = useService('debounceInput');
  const elementRegistry = useService('elementRegistry')


  const getValue = () => {
    return element.businessObject.participant || '';
  };

  const setValue = value => {
    return modeling.updateProperties(element, {
      participant: value
    });
  };

  const [ spells, setSpells ] = useState([]);

  useEffect(() => {
    function fetchSpells() {
      /*const hardcodedParticipants = ['Alice', 'Bob', 'Charlie', 'Dave'];
      setSpells(hardcodedParticipants);*/
      const participantElements = elementRegistry.filter(el => is(el, 'bpmn:Participant'));
      const participantNames = participantElements.map(el => {
        const bo = getBusinessObject(el);
        return bo.name || 'Unnamed Participant';
      });
      setSpells(participantNames);
    }

    fetchSpells();
  }, [ setSpells ]);

  const getOptions = () => {
    return [
      { label: '<none>', value: undefined },
      ...spells.map(spell => ({
        label: spell,
        value: spell
      }))
    ];
  };

  return html`<${SelectEntry}
    id=${ id }
    element=${ element }
    description=${ translate('Apply a black magic spell') }
    label=${ translate('Spell') }
    getValue=${ getValue }
    setValue=${ setValue }
    getOptions=${ getOptions }
    debounce=${ debounce }
  />`;
}