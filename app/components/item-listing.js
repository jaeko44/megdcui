import Ember from 'ember';
import PagingActions from 'megd/mixins/paging-actions';
export default Ember.Component.extend(PagingActions, {
  classNames: ['panel', 'panel-primary']
});
