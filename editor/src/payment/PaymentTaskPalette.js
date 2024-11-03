const SUITABILITY_SCORE_HIGH = 100,
      SUITABILITY_SCORE_AVERAGE = 50,
      SUITABILITY_SCORE_LOW = 25;

export default class PaymentTaskPalette {
  constructor(bpmnFactory, create, elementFactory, palette, translate) {
    this.bpmnFactory = bpmnFactory;
    this.create = create;
    this.elementFactory = elementFactory;
    this.translate = translate;

    palette.registerProvider(this);
  }

  getPaletteEntries(element) {
    const {
      bpmnFactory,
      create,
      elementFactory,
      translate
    } = this;

    function createTask(suitabilityScore) {
      return function(event) {
        const businessObject = bpmnFactory.create('bpmn:Task');

        //businessObject.suitable = suitabilityScore;
        businessObject.$attrs.type = 'paymentTask';

        const shape = elementFactory.createShape({
          type: 'bpmn:Task',
          businessObject: businessObject
        });

        create.start(event, shape);
      };
    }

    return {
      'create.payment-task': {
        group: 'activity',
        className: 'bpmn-icon-task green',
        title: translate('Create Payment Task'),
        action: {
          dragstart: createTask(SUITABILITY_SCORE_LOW),
          click: createTask(SUITABILITY_SCORE_LOW)
        }
      },
    };
  }
}

PaymentTaskPalette.$inject = [
  'bpmnFactory',
  'create',
  'elementFactory',
  'palette',
  'translate'
];