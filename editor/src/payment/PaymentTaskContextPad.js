export default class PaymentTaskContextPad {
    constructor(bpmnFactory, config, contextPad, create, elementFactory, injector, translate) {
      this.bpmnFactory = bpmnFactory;
      this.create = create;
      this.elementFactory = elementFactory;
      this.translate = translate;
  
      if (config.autoPlace !== false) {
        this.autoPlace = injector.get('autoPlace', false);
      }
  
      contextPad.registerProvider(this);
    }
  
    getContextPadEntries(element) {
      const {
        autoPlace,
        bpmnFactory,
        create,
        elementFactory,
        translate
      } = this;
  
      function appendPaymentTask() {
        return function(event, element) {
          if (autoPlace) {
            const businessObject = bpmnFactory.create('bpmn:Task');
  
            businessObject.$attrs.type = 'paymentTask';
  
            const shape = elementFactory.createShape({
              type: 'bpmn:Task',
              businessObject: businessObject
            });
  
            autoPlace.append(element, shape);
          } else {
            appendPaymentTaskStart(event, element);
          }
        };
      }
  
      function appendPaymentTaskStart() {
        return function(event) {
          const businessObject = bpmnFactory.create('bpmn:Task');
  
          businessObject.$attrs.type = 'paymentTask';
  
          const shape = elementFactory.createShape({
            type: 'bpmn:Task',
            businessObject: businessObject
          });
  
          create.start(event, shape, element);
        };
      }
  
      return {
        'append.low-task': {
          group: 'model',
          className: 'bpmn-icon-task green',
          title: translate('Append Payment Task'),
          action: {
            click: appendPaymentTask(),
            dragstart: appendPaymentTaskStart()
          }
        }
      };
    }
  }
  
  PaymentTaskContextPad.$inject = [
    'bpmnFactory',
    'config',
    'contextPad',
    'create',
    'elementFactory',
    'injector',
    'translate'
  ];