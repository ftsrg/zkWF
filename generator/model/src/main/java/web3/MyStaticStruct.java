package web3;

import org.web3j.abi.datatypes.*;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public abstract class MyStaticStruct extends StaticStruct {

    private final List<Class<Type>> itemTypes = new ArrayList<>();

    @SuppressWarnings("unchecked")
    public MyStaticStruct(List<Type> values) {
        super(values);
        for (Type value : values) {
            itemTypes.add((Class<Type>) value.getClass());
        }
    }

    @SafeVarargs
    public MyStaticStruct(Type... values) {
        this(Arrays.asList(values));
    }

    @Override
    public String getTypeAsString() {
        final StringBuilder type = new StringBuilder("(");
        for (int i = 0; i < itemTypes.size(); ++i) {
            final Class<Type> cls = itemTypes.get(i);
            if (StructType.class.isAssignableFrom(cls)) {
                type.append(getValue().get(i).getTypeAsString());
            }else if(StaticArray.class.isAssignableFrom(cls)) {
                type.append(getValue().get(i).getTypeAsString());
            } else {
                type.append(AbiTypes.getTypeAString(cls));
            }
            if (i < itemTypes.size() - 1) {
                type.append(",");
            }
        }
        type.append(")");
        return type.toString();
    }
}
